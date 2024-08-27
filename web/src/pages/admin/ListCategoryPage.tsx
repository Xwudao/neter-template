import {
  getApiGetSortDataCategory,
  getApiListAllCategory,
  postApiDeleteCategory,
  postApiUpdateSortCategory,
} from '@/api/categoryApi.ts';
import UploadBtn from '@/components/UploadBtn.tsx';
import { onError, onSuccess } from '@/core/callback.ts';
import { selectOptions } from '@/core/select_types.ts';
import SortByLoadBtn from '@/pages/admin/components/SortByLoadBtn.tsx';
import AdminContent from '@/pages/admin/layout/AdminContent.tsx';
import AdminToolbarTitle from '@/pages/admin/layout/AdminToolbarTitle.tsx';
import AdminWrapper from '@/pages/admin/layout/AdminWrapper.tsx';
import AddCategory from '@/pages/admin/popovers/AddCategory.tsx';
import EditCategory from '@/pages/admin/popovers/EditCategory.tsx';
import { adminPageRoute } from '@/router/routes.tsx';
import { arrayToTree } from '@/utils/tree.ts';
import { IconDelete, IconEdit, IconPlus, IconRefresh } from '@douyinfe/semi-icons';
import { Button, Popconfirm, Popover, Select, Table } from '@douyinfe/semi-ui';
import { ColumnProps } from '@douyinfe/semi-ui/lib/es/table';
import { useMutation, useQuery } from '@tanstack/react-query';
import { createRoute } from '@tanstack/react-router';

type IListCategoryPage = {};
const ListCategoryPage = (props: IListCategoryPage) => {
  console.log('ListCategoryPage render...');
  const [byOrder, setByOrder] = useState('desc');

  const { data, isLoading, isFetching, refetch } = useQuery({
    queryKey: ['list_category', byOrder],
    queryFn: () => getApiListAllCategory({ by_order: byOrder }),
  });

  const treeData = useMemo(() => {
    return arrayToTree(data?.data || []);
  }, [data?.data]);

  // use mutate for delete by id
  const { mutate: doDelete, isPending: deleting } = useMutation({
    mutationFn: postApiDeleteCategory,
  });

  const handleDelete = (id: number) => {
    doDelete(
      { id },
      {
        onSuccess: onSuccess('删除成功', () => {
          refetch();
        }),
        onError: onError(),
      },
    );
  };
  const { mutate: doSort, isPending: sorting } = useMutation({
    mutationFn: postApiUpdateSortCategory,
  });

  const handleSort = (ids: number[]) => {
    const orders = ids.map((_, index) => ids.length - index);
    doSort(
      { ids, orders },
      {
        onSuccess: onSuccess('更新成功', () => {
          refetch();
        }),
        onError: onError(),
      },
    );
  };

  const columns: ColumnProps[] = [
    {
      title: '分类名称',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: '分类描述',
      dataIndex: 'description',
      key: 'description',
    },
    {
      title: '权重排序',
      dataIndex: 'item_order',
      key: 'item_order',
    },
    {
      title: '操作',
      key: 'action',
      render: (text, record) => {
        return (
          <div className={`flex items-center gap2`}>
            <Popconfirm
              position={`bottomRight`}
              title={`确认删除吗`}
              content={`需要该分类下没有子分类且没有站点所属`}
              onConfirm={() => handleDelete(record.id)}>
              <Button type={`danger`} icon={<IconDelete />} size={`small`} loading={deleting} />
            </Popconfirm>
            <Popover
              showArrow
              position={`bottomRight`}
              trigger={`click`}
              content={
                <EditCategory onEdited={() => refetch()} category={record} categories={data?.data || []} />
              }>
              <Button icon={<IconEdit />} size={`small`} />
            </Popover>
          </div>
        );
      },
    },
  ];

  return (
    <>
      <AdminWrapper
        toolbar={
          <div className={`line-center justify-between flex-wrap`}>
            <AdminToolbarTitle className={`text-sm`}>分类管理</AdminToolbarTitle>
            <div className={`line-center gap-2 flex-wrap`}>
              <Select
                className={`w-24`}
                placeholder={`权重排序`}
                onChange={(v) => {
                  setByOrder(v as string);
                }}
                value={byOrder}
                optionList={selectOptions}
              />
              <UploadBtn />
              <SortByLoadBtn
                key={`load-cate`}
                loading={sorting}
                loadData={getApiGetSortDataCategory}
                onOk={(data) => handleSort(data.map((d) => d.id))}
              />
              {/*<SortCateBtn loading={sorting} onOk={(data) => handleSort(data.map((d) => d.id))} />*/}
              <Popover
                showArrow
                position={`bottomRight`}
                trigger={`click`}
                content={<AddCategory onAdded={() => refetch()} categories={data?.data || []} />}>
                <Button icon={<IconPlus />} />
              </Popover>
              <Button icon={<IconRefresh />} onClick={() => refetch()} loading={isFetching} />
            </div>
          </div>
        }>
        <AdminContent>
          <Table
            rowKey={`id`}
            columns={columns}
            dataSource={treeData}
            defaultExpandAllRows={treeData.length != 0}
            loading={isLoading}
            pagination={false}
          />
        </AdminContent>
      </AdminWrapper>
    </>
  );
};

const ListCategoryPageRoute = createRoute({
  getParentRoute: () => adminPageRoute,
  component: () => <ListCategoryPage />,
  path: '/list_category',
});

export default ListCategoryPageRoute;

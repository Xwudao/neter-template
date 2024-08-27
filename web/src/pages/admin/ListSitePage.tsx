import { getApiListAllCategory } from '@/api/categoryApi.ts';
import {
  getApiGetSortDataSite,
  getApiListSites,
  postApiDeleteSite,
  postApiUpdateSortSite,
  Site,
} from '@/api/siteApi.ts';
import { onError, onSuccess } from '@/core/callback.ts';
import { selectOptions } from '@/core/select_types.ts';
import SortByLoadBtn from '@/pages/admin/components/SortByLoadBtn.tsx';

import AdminContent from '@/pages/admin/layout/AdminContent.tsx';
import AdminToolbarTitle from '@/pages/admin/layout/AdminToolbarTitle.tsx';
import AdminWrapper from '@/pages/admin/layout/AdminWrapper.tsx';
import TableText from '@/pages/admin/layout/TableText.tsx';
import AddBulkSite from '@/pages/admin/popovers/AddBulkSite.tsx';

import AddSite from '@/pages/admin/popovers/AddSite.tsx';
import EditSite from '@/pages/admin/popovers/EditSite.tsx';
import { adminPageRoute } from '@/router/routes.tsx';
import { IconDelete, IconEdit, IconFixedStroked, IconPlus, IconRefresh } from '@douyinfe/semi-icons';
import { Button, Popconfirm, Popover, Select, Table } from '@douyinfe/semi-ui';
import { ColumnProps } from '@douyinfe/semi-ui/lib/es/table';
import { useMutation, useQuery } from '@tanstack/react-query';
import { createRoute } from '@tanstack/react-router';
import { SetStateAction } from 'react';

type IListSitePage = {};
const ListSitePage = (props: IListSitePage) => {
  console.log('ListSitePage render...');

  const [showBulk, setShowBulk] = useState(false);

  const [byOrder, setByOrder] = useState('desc');
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(10);
  const [categoryID, setCategoryID] = useState(0);

  const { data: categories } = useQuery({
    queryKey: ['list_category'],
    queryFn: () => getApiListAllCategory({ by_order: 'desc' }),
  });
  const {
    data: sites,
    isLoading,
    refetch,
    isFetching,
  } = useQuery({
    queryKey: ['list-sites', categoryID, page, size, byOrder],
    queryFn: () => getApiListSites({ page, size, category_id: categoryID, by_order: byOrder }),
  });

  // use mutate for delete by id
  const { mutate: doDelete, isPending: deleting } = useMutation({
    mutationFn: postApiDeleteSite,
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
    mutationFn: postApiUpdateSortSite,
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
      title: '站点名称',
      dataIndex: 'name',
      key: 'name',
      width: 300,
      render: (text: string, record: Site) => {
        return (
          <div className={`flex items-center gap2`}>
            <img src={record.icon} alt={record.name} width={20} height={20} className={`rd-full`} />
            <span>{record.name}</span>
          </div>
        );
      },
    },
    {
      title: '权重',
      dataIndex: 'item_order',
      key: 'item_order',
    },
    {
      title: '地址',
      dataIndex: 'url',
      key: 'url',
    },
    {
      title: '描述',
      dataIndex: 'description',
      key: 'description',
      width: 300,
      render(text: string) {
        return <TableText width={300} text={text} />;
      },
    },
    {
      title: '关键词',
      dataIndex: 'keywords',
      key: 'keywords',
      width: 300,
      render(text: string[]) {
        return <TableText width={300} text={text?.join(',')} />;
      },
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
              content={`这将同时删除站点所属分类关系`}
              onConfirm={() => handleDelete(record.id)}>
              <Button type={`danger`} icon={<IconDelete />} size={`small`} loading={deleting} />
            </Popconfirm>

            <Popover
              showArrow
              position={`bottomRight`}
              trigger={`click`}
              content={<EditSite onEdited={() => refetch()} data={record} categories={categories?.data || []} />}>
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
            <AdminToolbarTitle className={`text-sm`}>站点列表</AdminToolbarTitle>
            <div className={`line-center gap-2 flex-wrap`}>
              <Popover
                visible={showBulk}
                showArrow
                position={`bottomRight`}
                trigger={`custom`}
                content={<AddBulkSite onAdded={() => refetch()} categories={categories?.data || []} />}>
                <Button onClick={() => setShowBulk(!showBulk)} icon={<IconFixedStroked />} />
              </Popover>
              <Popover
                showArrow
                position={`bottomRight`}
                trigger={`click`}
                content={<AddSite onAdded={() => refetch()} categories={categories?.data || []} />}>
                <Button icon={<IconPlus />} />
              </Popover>
              <Select
                className={`w-24`}
                placeholder={`权重排序`}
                onChange={(v) => {
                  setPage(1);
                  setByOrder(v as string);
                }}
                value={byOrder}
                optionList={selectOptions}
              />
              <Select
                className={`w-40`}
                position={`bottomRight`}
                placeholder={`选择分类`}
                showClear
                value={categoryID === 0 ? undefined : categoryID}
                onChange={(v) => {
                  setPage(1);
                  setCategoryID(v as number);
                }}
                optionList={categories?.data?.map((item) => ({
                  value: item.id,
                  label: item.name,
                }))}
              />
              <SortByLoadBtn
                key={`load-site`}
                disabled={categoryID === 0}
                loadData={() => getApiGetSortDataSite({ pid: categoryID })}
                onOk={(data) => handleSort(data.map((d) => d.id))}
                loading={sorting}
              />
              <Button icon={<IconRefresh />} onClick={() => refetch()} loading={isFetching} />
            </div>
          </div>
        }>
        <AdminContent>
          <Table
            rowKey={`id`}
            columns={columns}
            dataSource={sites?.data.list || []}
            loading={isLoading}
            pagination={{
              total: sites?.data.total || 0,
              currentPage: page,
              pageSize: size,
              pageSizeOpts: [10, 20, 30, 40, 50],
              showSizeChanger: true,
              onPageSizeChange: (size: number) => {
                setPage(1);
                setSize(size);
              },
              onChange: (p: SetStateAction<number>, s: SetStateAction<number>) => {
                setPage(p);
                setSize(s);
              },
            }}
          />
        </AdminContent>
      </AdminWrapper>
    </>
  );
};

const ListSitePageRoute = createRoute({
  getParentRoute: () => adminPageRoute,
  component: () => <ListSitePage />,
  path: '/list_site',
});

export default ListSitePageRoute;

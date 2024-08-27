import {
  DataList,
  getApiGetDataListSortData,
  getApiListDataListByKind,
  postApiDeleteDataList,
  postApiUpdateSortDataList,
} from '@/api/dataListApi.ts';
import { onError, onSuccess } from '@/core/callback.ts';
import { DataListKinds } from '@/core/datalist_types.ts';
import SortByLoadBtn, { DataType } from '@/pages/admin/components/SortByLoadBtn.tsx';
import AdminContent from '@/pages/admin/layout/AdminContent.tsx';
import AdminToolbarTitle from '@/pages/admin/layout/AdminToolbarTitle.tsx';
import AdminWrapper from '@/pages/admin/layout/AdminWrapper.tsx';
import AddDataList from '@/pages/admin/popovers/AddDataList.tsx';
import EditDataList from '@/pages/admin/popovers/EditDataList.tsx';
import { adminPageRoute } from '@/router/routes.tsx';
import { formatDate } from '@/utils/date.ts';
import { IconDelete, IconEdit, IconPlus, IconRefresh } from '@douyinfe/semi-icons';
import { Button, Popconfirm, Popover, Select, Table } from '@douyinfe/semi-ui';
import { ColumnProps } from '@douyinfe/semi-ui/lib/es/table';
import { useMutation, useQuery } from '@tanstack/react-query';
import { createRoute } from '@tanstack/react-router';
import { SetStateAction } from 'react';

const DataListPage = () => {
  console.log('DataListPage render...');

  const [kind, setKind] = useState(DataListKinds[0].value);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(10);

  const { data, isLoading, isFetching, refetch } = useQuery({
    queryKey: ['list-datalist', kind, page, size],
    queryFn: () => getApiListDataListByKind({ kind: kind, page, size }),
  });
  const { mutate: doDelete, isPending: deleting } = useMutation({
    mutationFn: postApiDeleteDataList,
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
    mutationFn: postApiUpdateSortDataList,
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
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
    },
    {
      title: '名称',
      dataIndex: 'label',
      key: 'label',
    },
    {
      title: '排序',
      dataIndex: 'item_order',
      key: 'item_order',
    },
    {
      title: '类别',
      dataIndex: 'kind',
      key: 'kind',
    },
    {
      title: '内容值',
      dataIndex: 'value',
      key: 'value',
    },
    {
      title: '创建',
      dataIndex: 'create_time',
      key: 'create_time',
      render(text) {
        return <span>{formatDate(text)}</span>;
      },
    },
    {
      title: '操作',
      key: 'action',
      render: (text, record: DataList) => {
        return (
          <div className={`flex items-center gap2`}>
            <Popconfirm
              position={`bottomRight`}
              title={`确认删除吗`}
              content={`删除操作不可逆`}
              onConfirm={() => handleDelete(record.id)}>
              <Button type={`danger`} icon={<IconDelete />} size={`small`} loading={deleting} />
            </Popconfirm>
            <Popover
              showArrow
              position={`bottomRight`}
              trigger={`click`}
              content={<EditDataList onEdited={() => refetch()} data={record} kind={record.kind} />}>
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
            <AdminToolbarTitle className={`text-sm`}>数据管理</AdminToolbarTitle>
            <div className={`line-center gap-2 flex-wrap`}>
              <Select
                value={kind}
                className={`w-40`}
                onChange={(value) => setKind(value as string)}
                optionList={DataListKinds}
                placeholder={`选择类型`}
              />
              <SortByLoadBtn
                key={`load-datalist`}
                disabled={!kind}
                loadData={() => {
                  return new Promise<{ data: DataType[] }>((resolve, reject) => {
                    getApiGetDataListSortData({ kind: kind })
                      .then((rtn) => {
                        const data = rtn.data.map((item) => ({
                          id: item.id,
                          ...JSON.parse(item.value),
                        }));
                        resolve({ data });
                      })
                      .catch((err) => reject(err));
                  });
                }}
                onOk={(data) => handleSort(data.map((d) => d.id))}
                loading={sorting}
              />
              <Popover showArrow position={`bottomRight`} trigger={`click`} content={<AddDataList />}>
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
            dataSource={data?.data.list || []}
            loading={isLoading}
            pagination={{
              total: data?.data.total || 0,
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

const DataListPageRoute = createRoute({
  getParentRoute: () => adminPageRoute,
  component: () => <DataListPage />,
  path: '/data_list',
});

export default DataListPageRoute;

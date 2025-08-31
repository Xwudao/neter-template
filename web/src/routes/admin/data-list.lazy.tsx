import { DataList, postApiDeleteDataList } from '@/api/dataListApi';
import { getAdminApiListDataListByKind } from '@/api/dataListApi copy';
import AddDataListBtn from '@/components/admin/buttons/AddDataListBtn';
import DataViewBtn from '@/components/admin/buttons/DataViewBtn';
import ReverseStorePropBtn from '@/components/admin/buttons/ReverseStorePropBtn';
import AdminToolbarTitle from '@/components/admin/layout/AdminToolbarTitle';
import AdminWrapper from '@/components/admin/layout/AdminWrapper';
import EditDataList from '@/components/admin/surface/EditDataList';
import AppIcon from '@/components/AppIcon';
import { onSuccess } from '@/core/callback';
import KindTypes, { KindTypeLabels } from '@/core/kind_types';
import useDataListUpdate from '@/hooks/api/useDataListUpdate';
import { formatDate } from '@/utils/date.ts';
import { extractTitleName, hasProp } from '@/utils/name';
import { hasEnable } from '@/utils/object';
import { IconDelete, IconEdit, IconRefresh } from '@douyinfe/semi-icons';
import { Button, Popconfirm, Popover, Select, Table, Tag } from '@douyinfe/semi-ui';
import { ColumnProps } from '@douyinfe/semi-ui/lib/es/table';
import { useMutation, useQuery } from '@tanstack/react-query';
import { createLazyFileRoute } from '@tanstack/react-router';
import { useState } from 'react';
import { useLocalStorage } from 'react-use';

const DataListComponent = () => {
  const [page, setPage] = useState(1);
  const [memSize, setMemSize] = useLocalStorage('data-list-size', 10);
  const [size, setSize] = useState(memSize || 10);
  const [kind, setKind] = useState('');
  const [editingItem, setEditingItem] = useState<DataList | null>(null);
  const [updatingID, setUpdatingID] = useState<number | null>(null);

  const { updateProp, updating } = useDataListUpdate();

  const {
    data: dataLists,
    isLoading,
    isFetching,
    refetch,
  } = useQuery({
    queryKey: ['list-data-list', kind, page, size],
    queryFn: () =>
      getAdminApiListDataListByKind({
        kind: kind ? kind : '',
        page,
        size,
      }),
  });

  const { mutate: doDelete, isPending: deleting } = useMutation({
    mutationKey: ['delete-data-list'],
    mutationFn: postApiDeleteDataList,
  });

  const handleDelete = (id: number) => {
    doDelete(
      { id },
      {
        onSuccess: onSuccess(true, () => {
          refetch();
        }),
      },
    );
  };

  const handleEdit = (item: DataList) => {
    setEditingItem(item);
  };

  const columns: ColumnProps[] = [
    {
      title: 'ID',
      dataIndex: 'id',
      width: 80,
    },
    {
      title: '标签',
      dataIndex: 'label',
      width: 120,
    },
    {
      title: '类型',
      dataIndex: 'kind',
      width: 120,
    },
    {
      title: '名称',
      width: 120,
      render: (_, row: DataList) => {
        const name = extractTitleName(row.value);
        return <span>{name || '-'}</span>;
      },
    },
    {
      title: '状态',
      width: 100,
      render: (_, row: DataList) => {
        const enabled = hasEnable(row.value); //understand if the value has enable property
        return (
          <Tag
            color={enabled ? 'green' : 'red'}
            className={'cp'}
            onClick={() => {
              setUpdatingID(row.id);
              updateProp(
                row,
                'enable',
                (currentValue) => !currentValue,
                () => {
                  setUpdatingID(null);
                  refetch();
                },
              );
            }}>
            <AppIcon
              icon={'i-icon-park-outline-reverse-operation-out'}
              size={'small'}
              className={'gap1'}
              spin={updating && updatingID === row.id}
              text={enabled ? '启用' : '禁用'}
            />
            {/* {enabled ? '启用' : '禁用'} */}
          </Tag>
        );
      },
    },
    // {
    //   title: '键',
    //   dataIndex: 'key',
    //   width: 150,
    // },
    {
      title: '值',
      dataIndex: 'value',
      width: 200,
      render: (text: string) => (
        <div className="max-w-48 truncate" title={text}>
          {text}
        </div>
      ),
    },
    {
      title: '创建时间',
      dataIndex: 'create_time',
      width: 180,
      render: (text) => formatDate(text),
    },
    {
      title: '更新时间',
      dataIndex: 'update_time',
      width: 180,
      render: (text) => formatDate(text),
    },
    {
      title: '操作',
      width: 120,
      fixed: 'right',
      render: (_, row: DataList) => (
        <div className="flex gap-2">
          <Popover
            content={
              <EditDataList
                className="p-4 w-70"
                dataList={row}
                onClose={() => setEditingItem(null)}
                onSuccess={() => {
                  refetch();
                  setEditingItem(null);
                }}
              />
            }
            trigger="click"
            visible={editingItem?.id === row.id}
            onVisibleChange={(visible) => {
              if (!visible) setEditingItem(null);
            }}>
            <Button size="small" icon={<IconEdit />} onClick={() => handleEdit(row)} />
          </Popover>
          {hasProp(row.value, 'enable') && (
            <ReverseStorePropBtn
              size="small"
              datalist={row}
              prop={`enable`}
              onRefresh={() => refetch()}
              name={`启用状态`}
            />
          )}

          <DataViewBtn kind_type={row.kind as KindTypes} disabled={row.kind !== KindTypes.BanUA} />

          <Popconfirm title="确认删除" content="删除后无法恢复，确定要删除吗？" onConfirm={() => handleDelete(row.id)}>
            <Button size="small" type="danger" icon={<IconDelete />} loading={deleting} />
          </Popconfirm>
        </div>
      ),
    },
  ];

  return (
    <>
      <AdminWrapper
        toolbar={
          <div className="line-center justify-between flex-wrap">
            <AdminToolbarTitle className="text-sm">数据管理</AdminToolbarTitle>
            <div className="line-center gap-2 flex-wrap">
              <Select
                placeholder="选择类型"
                value={kind}
                optionList={Object.keys(KindTypeLabels).map((key) => ({
                  value: key,
                  label: KindTypeLabels[key as keyof typeof KindTypeLabels],
                }))}
                onChange={(value) => {
                  setPage(1);
                  setKind(value as string);
                }}
                style={{ width: 150 }}
                showClear
              />
              <AddDataListBtn />
              <Button loading={isFetching} icon={<IconRefresh />} onClick={() => refetch()} />
            </div>
          </div>
        }>
        <Table
          className="w-full"
          dataSource={dataLists?.data?.list || []}
          columns={columns}
          rowKey="id"
          loading={isLoading}
          // rowSelection={{
          //   onChange: (selectedRowKeys) => {
          //     setSelectedIDs(selectedRowKeys as number[]);
          //   },
          // }}
          pagination={{
            total: dataLists?.data?.total || 0,
            pageSize: size,
            currentPage: page,
            onPageChange: (currentPage) => setPage(currentPage),
            pageSizeOpts: [10, 20, 30, 40, 50],
            showSizeChanger: true,
            onPageSizeChange: (size) => {
              setPage(1);
              setSize(size);
              setMemSize(size);
            },
          }}
        />
      </AdminWrapper>
    </>
  );
};

export const Route = createLazyFileRoute('/admin/data-list')({
  component: DataListComponent,
});

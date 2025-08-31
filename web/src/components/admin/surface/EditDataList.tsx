import { Button, Form, Toast } from '@douyinfe/semi-ui';
import { postApiUpdateDataList, PostUpdateDataListReq, DataList } from '@/api/dataListApi';
import { useMutation } from '@tanstack/react-query';
import { OK_CODE } from '@/api/types';
import DataListForm from '@/components/admin/forms/DataListForm';
import PopTitle from '@/components/PopTitle';
import clsx from 'clsx';

interface Props {
  className?: string;
  dataList: DataList;
  onClose?: () => void;
  onSuccess?: () => void;
}

function EditDataList(props: Props) {
  const { dataList, onSuccess, className, onClose } = props;

  const updateDataListMutation = useMutation({
    mutationKey: ['updateDataList'],
    mutationFn: postApiUpdateDataList,
    onSuccess: (res) => {
      if (res.code === OK_CODE) {
        Toast.success('更新成功');
        onSuccess?.();
      } else {
        Toast.error(res.msg || '更新失败');
      }
    },
    onError: () => {
      Toast.error('更新失败');
    },
  });

  const handleSubmit = async (values: any) => {
    const payload: PostUpdateDataListReq = {
      id: dataList.id,
      key: values.key,
      value: values.value,
    };
    updateDataListMutation.mutate(payload);
  };

  const initialValues = {
    label: dataList.label,
    kind: dataList.kind,
    key: dataList.key,
    value: dataList.value,
  };

  return (
    <div className={clsx(className)}>
      <Form onSubmit={handleSubmit} initValues={initialValues}>
        <PopTitle title="编辑数据" onClose={onClose} />
        <DataListForm isEdit />
        <div className="flex justify-end mt-4">
          <Button type="primary" htmlType="submit" className={'w-full'} loading={updateDataListMutation.isPending}>
            更新
          </Button>
        </div>
      </Form>
    </div>
  );
}

export default EditDataList;

import { postApiCreateDataList, PostCreateDataListReq } from '@/api/dataListApi';
import { OK_CODE } from '@/api/types';
import DataListForm from '@/components/admin/forms/DataListForm';
import PopTitle from '@/components/PopTitle';
import { KindTypeLabels } from '@/core/kind_types';
import { Button, Form, Toast } from '@douyinfe/semi-ui';
import { useMutation } from '@tanstack/react-query';
import clsx from 'clsx';

interface Props {
  className?: string;
  showTitle?: boolean;
  onClose?: () => void;
  onSuccess?: () => void;
}

function AddDataList(props: Props) {
  const { onSuccess, className, onClose, showTitle } = props;

  const createDataListMutation = useMutation({
    mutationKey: ['createDataList'],
    mutationFn: postApiCreateDataList,
    onSuccess: (res) => {
      if (res.code === OK_CODE) {
        Toast.success('创建成功');
        onSuccess?.();
      } else {
        Toast.error(res.msg || '创建失败');
      }
    },
    onError: () => {
      Toast.error('创建失败');
    },
  });

  const handleSubmit = async (values: PostCreateDataListReq) => {
    // build: label,
    values.label = KindTypeLabels[values.kind as keyof typeof KindTypeLabels] || '';
    createDataListMutation.mutate(values);
  };

  return (
    <div className={clsx(className)}>
      <Form onSubmit={handleSubmit}>
        {showTitle && <PopTitle title="创建列表" onClose={onClose} />}
        <DataListForm />
        <div className="flex justify-end mt-4">
          <Button type="primary" htmlType="submit" className={'w-full'} loading={createDataListMutation.isPending}>
            创建
          </Button>
        </div>
      </Form>
    </div>
  );
}

export default AddDataList;

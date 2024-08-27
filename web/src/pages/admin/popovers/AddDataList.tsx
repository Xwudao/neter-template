import { postApiCreateDataList } from '@/api/dataListApi.ts';
import { onError, onSuccess } from '@/core/callback.ts';
import DataListWrapper from '@/pages/admin/datalist/DataListWrapper.tsx';
import { Button } from '@douyinfe/semi-ui';
import { useMutation } from '@tanstack/react-query';
import { FC, PropsWithChildren } from 'react';

type IAddDataList = {
  onAdded?: () => void;
};
const AddDataList: FC<PropsWithChildren<IAddDataList>> = ({ onAdded }) => {
  console.log('addDataList render...');

  const { mutate: doAdd, isPending } = useMutation({
    mutationFn: postApiCreateDataList,
  });
  const handleAdd = (values: any) => {
    doAdd(
      { ...values },
      {
        onSuccess: onSuccess('添加成功', () => onAdded?.()),
        onError: onError(),
      },
    );
  };

  return (
    <div className={`p2 w-60`}>
      <DataListWrapper title={`添加数据`} onSubmit={handleAdd}>
        <Button block loading={isPending} htmlType={`submit`}>
          提交
        </Button>
      </DataListWrapper>
    </div>
  );
};

export default AddDataList;

import { DataList, postApiUpdateDataList } from '@/api/dataListApi.ts';
import { onError, onSuccess } from '@/core/callback.ts';
import DataListWrapper from '@/pages/admin/datalist/DataListWrapper.tsx';
import { Button } from '@douyinfe/semi-ui';
import { useMutation } from '@tanstack/react-query';
import { FC, PropsWithChildren } from 'react';
import { md5 } from 'js-md5';

type IEditDataList = {
  data: DataList;
  kind: string;
  onEdited?: () => void;
};
const EditDataList: FC<PropsWithChildren<IEditDataList>> = ({ data, kind, onEdited }) => {
  console.log('editDataList render...');
  const { mutate: doEdit, isPending } = useMutation({
    mutationFn: postApiUpdateDataList,
  });
  const handleEdit = (values: any) => {
    doEdit(
      {
        ...values,
        key: md5(JSON.stringify(values)),
        id: data.id,
      },
      {
        onSuccess: onSuccess(true, () => onEdited?.()),
        onError: onError(),
      },
    );
  };
  return (
    <div className={`p2`}>
      <DataListWrapper
        title={`修改数据`}
        onSubmit={handleEdit}
        initValues={JSON.parse(data.value)}
        defaultKind={kind}>
        <Button block loading={isPending} htmlType={`submit`}>
          提交
        </Button>
      </DataListWrapper>
    </div>
  );
};

export default EditDataList;

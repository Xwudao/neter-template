import DataLinkForm from '@/pages/admin/datalist/DataLinkForm.tsx';
import { FC, PropsWithChildren } from 'react';

type IDataListForm = {
  kind: string;
};
const DataListForm: FC<PropsWithChildren<IDataListForm>> = ({ kind }) => {
  console.log('dataListForm render...');

  // const formApi = useFormApi();
  // const formState = useFormState();

  return <>{kind === 'friend_link' && <DataLinkForm />}</>;
};

export default DataListForm;

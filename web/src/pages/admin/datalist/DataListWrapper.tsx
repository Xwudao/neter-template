import { DataListKinds } from '@/core/datalist_types.ts';
import DataLinkForm from '@/pages/admin/datalist/DataLinkForm.tsx';
import { Form } from '@douyinfe/semi-ui';
import classnames from 'classnames';
import { md5 } from 'js-md5';
import { FC, PropsWithChildren } from 'react';
const { Select } = Form;

type IDataListWrapper = {
  title: string;
  defaultKind?: string;
  className?: string;
  initValues?: Record<string, any>;
  onSubmit: (values: Record<string, any>) => void;
};
const DataListWrapper: FC<PropsWithChildren<IDataListWrapper>> = ({
  defaultKind,
  title,
  children,
  className = '',
  initValues,
  onSubmit,
}) => {
  console.log('dataListWrapper render...');

  const [kind, setKind] = useState<string | undefined>(defaultKind);

  return (
    <div>
      <h5 className={`fw-bold text-base mb4`}>{title}</h5>

      <Form
        initValues={initValues}
        className={classnames(className)}
        onSubmit={(data) => {
          onSubmit({
            ...data,
            label: DataListKinds.find((item) => (item.value = data['kind']))?.label,
            key: md5(JSON.stringify(data)),
            value: JSON.stringify(data),
          });
        }}>
        {!defaultKind && (
          <Select
            field={`kind`}
            label={`类型`}
            onChange={(value) => setKind(value as string)}
            optionList={DataListKinds}
            className={`w-full`}
            placeholder={`选择类型`}
            rules={[{ required: true, message: '必填' }]}
          />
        )}

        {['friend_link', 'top_link'].includes(kind!) && <DataLinkForm />}
        {children}
      </Form>
    </div>
  );
};

export default DataListWrapper;

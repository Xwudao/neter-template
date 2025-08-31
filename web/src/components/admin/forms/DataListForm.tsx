import DataKindsBuilder from '@/components/others/DataKindsBuilder';
import KindTypes, { KindTypeLabels } from '@/core/kind_types';
import { Form, useFormApi, useFormState } from '@douyinfe/semi-ui';

interface Props {
  isEdit?: boolean;
}

const { TextArea, Select } = Form;

function DataListForm(props: Props) {
  const { isEdit = false } = props;

  const formApi = useFormApi();
  const formState = useFormState();

  return (
    <>
      <Select
        field={'kind'}
        label="类型"
        disabled={isEdit}
        className={'w-full'}
        placeholder="请选择类型"
        rules={[{ required: true, message: '类型不能为空' }]}
        optionList={Object.entries(KindTypeLabels).map(([value, label]) => ({ value, label }))}
        onChange={(value) => {
          formApi.setValue('label', KindTypeLabels[value as KindTypes] || '');
        }}
      />

      {formState.values.kind && (
        <DataKindsBuilder
          kind={formState.values.kind}
          onKeyChange={(key) => {
            formApi.setValue('key', key);
          }}
          value={formState.values.value}
          onChange={(value) => {
            formApi.setValue('value', value);
          }}
        />
      )}

      <TextArea
        field="value"
        label="原始值"
        disabled
        placeholder="JSON格式的配置值"
        autosize={{ minRows: 3, maxRows: 8 }}
        rules={[{ required: true, message: '值不能为空' }]}
        // style={{ display: 'none' }} // Hide raw value field since we use DataKindsBuilder
      />
    </>
  );
}

export default DataListForm;

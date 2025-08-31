import { Form } from '@douyinfe/semi-ui';
import { md5 } from 'js-md5';
import { useEffect, useState } from 'react';

const { Input, Checkbox } = Form;

interface BackupLinkData {
  name: string;
  link: string;
  open_blank: boolean;
  enable: boolean;
}

interface Props {
  value?: string;
  onChange?: (value: string) => void;
  onKeyChange?: (key: string) => void;
}

function BackupLinkForm({ value, onChange, onKeyChange }: Props) {
  const [formData, setFormData] = useState<BackupLinkData>(JSON.parse(value || '{}') as BackupLinkData);

  useEffect(() => {
    if (value) {
      try {
        const parsed = JSON.parse(value) as BackupLinkData;
        setFormData(parsed);
      } catch {
        // Invalid JSON, keep default values
      }
    }
  }, [value]);

  const handleFormChange = (values: BackupLinkData) => {
    setFormData(values);
    onChange?.(JSON.stringify(values));
  };

  return (
    <Form
      initValues={formData}
      onSubmit={(e) => {
        e.preventDefault();
      }}
      onValueChange={(val) => {
        handleFormChange(val as BackupLinkData);
        onKeyChange?.(md5(`backup_link-${val.name || ''}`));
      }}
      className="flex flex-col gap2"
      render={() => {
        return (
          <>
            <Input
              field="name"
              label="名称"
              placeholder="请输入备份链接名称"
              rules={[{ required: true, message: '名称不能为空' }]}
            />
            <Input
              field="link"
              label="链接"
              placeholder="请输入备份链接地址"
              rules={[{ required: true, message: '链接不能为空' }]}
            />
            <div className="flex items-center gap3">
              <Checkbox field="open_blank" label="新窗口打开" className={'flex-1'}>
                在新标签打开
              </Checkbox>
              <Checkbox field="enable" label="启用" className={'flex-1'}>
                启用该备份链接
              </Checkbox>
            </div>
          </>
        );
      }}
    />
  );
}

export default BackupLinkForm;

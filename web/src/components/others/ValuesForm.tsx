import AppIcon from '@/components/AppIcon';
import { IconPlus } from '@douyinfe/semi-icons';
import { ArrayField, Button, Form } from '@douyinfe/semi-ui';
import { md5 } from 'js-md5';

const { Input, Checkbox } = Form;

interface Props {
  field: string;
  initValue?: string[];
  removeEmpty?: boolean; // 是否移除空字符串
  handleFormChange?: (values: string[]) => void;
  onKeyChange?: (key: string) => void;
}

function ValuesForm({ field, initValue, removeEmpty = true, handleFormChange, onKeyChange }: Props) {
  console.log('🚀 ~ ValuesForm ~ initValue:', initValue);
  return (
    <Form
      initValues={{ [field]: initValue || [] }}
      onSubmit={(e) => {
        e.preventDefault();
      }}
      onValueChange={(val) => {
        // 如果需要移除空字符串项
        if (removeEmpty) {
          const values = (val[field] as string[]).filter((item) => item.trim() !== '');
          val[field] = values;
        }
        handleFormChange?.(val as string[]);
        onKeyChange?.(md5(`values-${val[field]?.join(',') || ''}`));
      }}>
      <Checkbox field="enable" label="启用">
        是否启用该配置
      </Checkbox>

      <ArrayField field={field}>
        {({ add, arrayFields }) => (
          <div className="flex flex-col gap-2">
            <div className="flex justify-between items-center">
              <span className="font-medium">数组内容</span>
              <Button icon={<IconPlus />} size="small" onClick={add} theme="borderless">
                添加
              </Button>
            </div>

            <div className={'max-h-56 overflow-y-auto'}>
              {arrayFields.map(({ field: itemField, key, remove }, index) => (
                <div key={key} className="flex gap-2 items-center w-full">
                  <Input
                    noLabel
                    field={itemField}
                    placeholder={`请输入第${index + 1}项`}
                    className="flex-1"
                    fieldClassName={'w-full'}
                  />
                  <Button
                    icon={<AppIcon icon={'i-material-symbols-delete-outline'} />}
                    type="danger"
                    theme="borderless"
                    size="default"
                    onClick={remove}
                  />
                </div>
              ))}
            </div>
            {arrayFields.length === 0 && <div className="">暂无数据，点击添加按钮添加项目</div>}
          </div>
        )}
      </ArrayField>
    </Form>
  );
}

export default ValuesForm;

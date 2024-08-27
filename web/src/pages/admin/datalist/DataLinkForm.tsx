import { Form } from '@douyinfe/semi-ui';
import { FC, PropsWithChildren } from 'react';

const { Checkbox, Input } = Form;

type IFriendLinkForm = {};
const DataLinkForm: FC<PropsWithChildren<IFriendLinkForm>> = (props) => {
  console.log('friendLinkForm render...');

  return (
    <>
      <Input field={`name`} label={`名称`} placeholder={`请输入名称`} rules={[{ required: true, message: '必填' }]} />
      <Input field={`link`} label={`链接`} placeholder={`请输入链接`} rules={[{ required: true, message: '必填' }]} />
      <div className={`flex gap2 justify-between`}>
        <Checkbox label={`打开`} field={`open_blank`} rules={[{ required: true, message: '必填' }]}>
          新窗口打开
        </Checkbox>
        <Checkbox label={`启用`} field={`enable`} rules={[{ required: true, message: '必填' }]}>
          启用后显示
        </Checkbox>
      </div>
    </>
  );
};

export default DataLinkForm;

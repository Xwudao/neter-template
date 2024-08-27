import { Form } from '@douyinfe/semi-ui';
import classnames from 'classnames';
import { FC, PropsWithChildren } from 'react';

const { TextArea, Input, Select } = Form;

type ISiteInfoForm = { className?: string };
const SiteInfoForm: FC<PropsWithChildren<ISiteInfoForm>> = ({ className }) => {
  console.log('siteInfoForm render...');
  return (
    <div className={classnames(className)}>
      <Input
        field={`site_name`}
        label={`站点名称`}
        placeholder={`请输入站点名称`}
        rules={[{ required: true, message: '必填' }]}
      />
      <Input
        field={`site_url`}
        label={`站点 URL`}
        placeholder={`请输入站点 URL`}
        rules={[{ required: true, message: '必填' }]}
      />
      <Input
        field={`site_title`}
        label={`站点标题`}
        placeholder={`请输入站点标题`}
        rules={[{ required: true, message: '必填' }]}
      />
      <Input
        field={`sub_title`}
        label={`副标题`}
        placeholder={`请输入副标题`}
        rules={[{ required: true, message: '必填' }]}
      />
      <Input
        field={`main_title`}
        label={`主页标题`}
        placeholder={`请输入主页标题`}
        rules={[{ required: true, message: '必填' }]}
      />
      <Input field={`site_image`} label={`站点图片`} placeholder={`主要用于Open Graph`} />
      <Input
        field={`site_logo`}
        label={`站点Logo`}
        placeholder={`请输入站点Logo`}
        rules={[{ required: true, message: '必填' }]}
      />

      <Select
        className={`w-full`}
        allowCreate={true}
        multiple={true}
        filter={true}
        field={`site_keywords`}
        label={`站点关键词`}
        placeholder={`请输入站点关键词，以逗号分隔`}
        rules={[{ required: true, message: '必填' }]}
      />
      <TextArea
        field={`site_desc`}
        label={`站点描述`}
        placeholder={`请输入站点描述`}
        rules={[{ required: true, message: '必填' }]}
      />
      <TextArea
        field={`disclaimer`}
        label={`免责声明`}
        placeholder={`请输入免责声明`}
        rules={[{ required: true, message: '必填' }]}
      />
      <TextArea
        field={`site_meta_script`}
        autosize={{ minRows: 1, maxRows: 10 }}
        label={`站点 Meta Script`}
        placeholder={`请输入站点 Meta Script`}
      />
    </div>
  );
};

export default SiteInfoForm;

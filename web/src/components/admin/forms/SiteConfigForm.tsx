import { Form, Typography, useFormApi } from '@douyinfe/semi-ui';

const { TextArea, Input, Select } = Form;

function SiteConfigForm() {
  // const {} = props

  const formApi = useFormApi();
  return (
    <>
      <Input
        field={`site_name`}
        label={`站点名称`}
        placeholder={`请输入站点名称`}
        rules={[{ required: true, message: '必填' }]}
      />

      <Input
        field={`site_title`}
        label={`站点标题`}
        placeholder={`请输入站点标题`}
        rules={[{ required: true, message: '必填' }]}
      />
      <Input
        field={`site_url`}
        label={`站点 URL`}
        placeholder={`请输入站点 URL`}
        rules={[{ required: true, message: '必填' }]}
        extraText={
          <Typography.Text
            size={'small'}
            link
            onClick={() => {
              formApi.setValue('site_url', window.location.origin);
            }}>
            使用当前网址
          </Typography.Text>
        }
      />
      {/* <Input field={`site_image`} label={`站点图片`} placeholder={`主要用于Open Graph`} /> */}

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
        field={`site_meta_script`}
        autosize={{ minRows: 2, maxRows: 10 }}
        label={`站点 Meta Script`}
        placeholder={`请输入站点 Meta Script`}
      />
    </>
  );
}

export default SiteConfigForm;

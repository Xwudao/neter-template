import ConfigFormWrapper from '@/components/admin/config/ConfigFormWrapper';
import SiteConfigForm from '@/components/admin/forms/SiteConfigForm';
import useAdminConfig from '@/provider/useAdminConfig';
import { IconSave } from '@douyinfe/semi-icons';
import { Button } from '@douyinfe/semi-ui';
import clsx from 'clsx';

interface Props {
  className?: string;
}

function SiteConfigTab(props: Props) {
  const { className } = props;
  const {
    config: { site_info },
    handleUpdateConfig,
    updating,
  } = useAdminConfig();
  return (
    <ConfigFormWrapper
      className={clsx(className)}
      initValues={site_info}
      onSubmit={(values) => {
        handleUpdateConfig('site_info', values, true);
        console.log('🚀 ~ <ConfigFormWrapperclassName={clsx ~ values:', values);
      }}>
      <SiteConfigForm />
      <Button htmlType={'submit'} loading={updating} type="primary" className={`mt-4`} icon={<IconSave />}>
        保存配置
      </Button>
    </ConfigFormWrapper>
  );
}

export default SiteConfigTab;

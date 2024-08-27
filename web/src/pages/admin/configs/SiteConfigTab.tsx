import ConfigFormWrapper from '@/pages/admin/configs/ConfigFormWrapper.tsx';
import SiteInfoForm from '@/pages/admin/configs/SiteInfoForm.tsx';
import { IconSave } from '@douyinfe/semi-icons';
import { Button } from '@douyinfe/semi-ui';
import { FC, PropsWithChildren } from 'react';

type ISiteConfigTab = {
  className?: string;
  value: string;
  isLoading: boolean;
  configKey: string;
  onSave?: (key: string, values: string) => void;
};
const SiteConfigTab: FC<PropsWithChildren<ISiteConfigTab>> = ({
  className = '',
  isLoading,
  value,
  configKey,
  onSave,
}) => {
  const obj = JSON.parse(value);
  console.log('siteConfigTab render...', obj);
  return (
    <div className={className}>
      <ConfigFormWrapper
        initValues={obj}
        onSubmit={(values) => {
          onSave?.(configKey, JSON.stringify(values));
        }}>
        <SiteInfoForm className={`grid grid-cols-2 gap4`} />
        <Button type={`primary`} htmlType={`submit`} icon={<IconSave />} loading={isLoading}>
          提交
        </Button>
      </ConfigFormWrapper>
    </div>
  );
};

export default SiteConfigTab;

import { getApiGenSitemap } from '@/api/siteConfigApi.ts';
import { onError, onSuccess } from '@/core/callback.ts';
import ConfigFormWrapper from '@/pages/admin/configs/ConfigFormWrapper.tsx';
import SEOForm from '@/pages/admin/configs/SEOForm.tsx';
import { IconSave } from '@douyinfe/semi-icons';
import { Button } from '@douyinfe/semi-ui';
import { useMutation } from '@tanstack/react-query';
import { FC, PropsWithChildren } from 'react';

type ISeoConfigTab = {
  className?: string;
  value: string;
  isLoading: boolean;
  configKey: string;
  onSave?: (key: string, values: string) => void;
};
const SeoConfigTab: FC<PropsWithChildren<ISeoConfigTab>> = ({
  className = '',
  isLoading,
  value,
  configKey,
  onSave,
}) => {
  console.log('seoConfigTab render...');
  const obj = JSON.parse(value);

  const { mutate: doGen, isPending: generating } = useMutation({
    mutationFn: getApiGenSitemap,
  });

  const handleGenSitemap = () => {
    doGen(undefined, {
      onSuccess: onSuccess('data'),
      onError: onError(),
    });
  };

  return (
    <div className={className}>
      <ConfigFormWrapper
        initValues={obj}
        onSubmit={(values) => {
          onSave?.(configKey, JSON.stringify(values));
        }}>
        <SEOForm className={`grid grid-cols-2 gap4`} />
        <div className={`flex gap2`}>
          <Button type={`primary`} htmlType={`submit`} icon={<IconSave />} loading={isLoading}>
            提交
          </Button>
          <Button type={`warning`} onClick={handleGenSitemap} loading={generating}>
            生成站点地图
          </Button>
        </div>
      </ConfigFormWrapper>
    </div>
  );
};

export default SeoConfigTab;

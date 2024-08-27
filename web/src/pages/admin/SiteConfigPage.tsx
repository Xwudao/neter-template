import { getApiGetSiteConfig, postApiUpdateSiteConfig } from '@/api/siteConfigApi.ts';
import Loading from '@/components/Loading.tsx';
import Show from '@/components/Show.tsx';
import UploadBtn from '@/components/UploadBtn.tsx';
import { onError, onSuccess } from '@/core/callback.ts';
import SeoConfigTab from '@/pages/admin/configs/SeoConfigTab.tsx';
import SiteConfigTab from '@/pages/admin/configs/SiteConfigTab.tsx';
import AdminContent from '@/pages/admin/layout/AdminContent.tsx';
import AdminToolbarTitle from '@/pages/admin/layout/AdminToolbarTitle.tsx';
import AdminWrapper from '@/pages/admin/layout/AdminWrapper.tsx';
import { adminPageRoute } from '@/router/routes.tsx';
import { IconRefresh } from '@douyinfe/semi-icons';
import { Button, TabPane, Tabs } from '@douyinfe/semi-ui';
import { useMutation, useQuery } from '@tanstack/react-query';
import { createRoute } from '@tanstack/react-router';

const SiteConfigPage = () => {
  console.log('SiteConfigPage render...');

  const { data, isLoading, refetch, isFetching } = useQuery({
    queryKey: ['list-all-configs'],
    queryFn: () => getApiGetSiteConfig(),
  });

  const { mutate: doCreate, isPending } = useMutation({
    mutationFn: postApiUpdateSiteConfig,
  });

  const handleSave = (key: string, value: string) => {
    doCreate(
      {
        name: key,
        config: value,
      },
      {
        onSuccess: onSuccess('保存成功', () => refetch()),
        onError: onError(),
      },
    );
  };

  return (
    <>
      <AdminWrapper
        toolbar={
          <div className={`line-center justify-between flex-wrap`}>
            <AdminToolbarTitle className={`text-sm`}>站点配置</AdminToolbarTitle>
            <div className={`line-center gap-2 flex-wrap`}>
              <UploadBtn prefix={`attachment`} />
              <Button icon={<IconRefresh />} onClick={() => refetch()} loading={isFetching} />
            </div>
          </div>
        }>
        <AdminContent>
          <Loading show={isLoading} className={`block-center my2`} />
          <Show show={!isLoading}>
            <Tabs type="line" defaultActiveKey={`site-config`} >
              <TabPane tab="站点配置" itemKey="site-config">
                <SiteConfigTab
                  isLoading={isLoading || isPending}
                  className={`w-160`}
                  value={data?.data?.site_info || '{}'}
                  configKey={`site_info`}
                  onSave={handleSave}
                />
              </TabPane>
              <TabPane tab="SEO配置" itemKey="seo-config">
                <SeoConfigTab
                  isLoading={isLoading || isPending}
                  className={`w-160`}
                  value={data?.data?.seo_config || '{}'}
                  configKey={`seo_config`}
                  onSave={handleSave}
                />
              </TabPane>
            </Tabs>
          </Show>
        </AdminContent>
      </AdminWrapper>
    </>
  );
};

const SiteConfigPageRoute = createRoute({
  getParentRoute: () => adminPageRoute,
  component: () => <SiteConfigPage />,
  path: '/site_config',
});

export default SiteConfigPageRoute;

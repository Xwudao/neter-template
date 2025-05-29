import AdminToolbarTitle from '@/components/admin/layout/AdminToolbarTitle';
import AdminWrapper from '@/components/admin/layout/AdminWrapper';
import SiteConfigTab from '@/components/admin/tabs/SiteConfigTab';
import { Tabs } from '@douyinfe/semi-ui';
import { createLazyFileRoute, useNavigate, useSearch } from '@tanstack/react-router';


// type ConfigSearch = z.infer<typeof configSearchSchema>;

const ConfigComponent = () => {
  const { tab } = useSearch({ from: '/admin' });
  const nav = useNavigate();

  return (
    <AdminWrapper
      toolbar={
        <div className={`line-center justify-between flex-wrap`}>
          <AdminToolbarTitle className={`text-sm`}>站点设置</AdminToolbarTitle>
          <div className={`line-center gap-2 flex-wrap`}></div>
        </div>
      }>
      <Tabs
        activeKey={tab}
        onChange={(t) => {
          nav({
            search: { tab: t },
          });
        }}>
        <Tabs.TabPane itemKey="config" tab="站点配置">
          <SiteConfigTab className={'max-w-full w-120'} />
        </Tabs.TabPane>
        <Tabs.TabPane itemKey="seo" tab="SEO配置">
          <SiteConfigTab className={'max-w-full w-120'} />
        </Tabs.TabPane>
      </Tabs>
    </AdminWrapper>
  );
};

export const Route = createLazyFileRoute('/admin/config')({
  component: () => <ConfigComponent />,
});

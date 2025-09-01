import appLogo from '@/assets/images/app.svg';
import Navs, { findKeyByPath, findPathByKey } from '@/components/admin/navs';
import AppIcon from '@/components/AppIcon';
import ContentLoading from '@/components/loading/ContentLoading';
import useTheme from '@/hooks/useTheme';
import AdminConfigProvider from '@/provider/AdminConfigProvider';
import useAuth from '@/provider/useAuth';
import { Avatar, Button, Divider, Dropdown, Layout, Nav, Toast } from '@douyinfe/semi-ui';
import { createFileRoute, Link, Outlet, redirect, useLocation, useNavigate } from '@tanstack/react-router';
import { Suspense, useMemo } from 'react';
import z from 'zod';
import MaterialSymbolsLogoutSharp from '~icons/material-symbols/logout-sharp';
import classes from '../styles.module.scss';

const { Header, Footer, Sider, Content } = Layout;

const configSearchSchema = z.object({
  tab: z.string().optional().default('config'),
});

const AdminLayout = () => {
  const nav = useNavigate();
  const { logout, user } = useAuth();
  const { toggleTheme, isDark } = useTheme();
  const loc = useLocation();

  const defaultKeys = useMemo(() => {
    const key = findKeyByPath(Navs, loc.pathname);
    if (key) return [key];
    return ['admin-index'];
  }, [loc]);

  return (
    <>
      <Layout className={classes.adminMain}>
        <Header className={classes.adminMainHeader}>
          <Nav mode={`horizontal`}>
            <Nav.Header>
              <img src={appLogo} alt={`Logo`} style={{ height: '20px', fontSize: 36 }} />
              <span className={'text-lg ml2'}>无道后台</span>
            </Nav.Header>
            <Nav.Footer>
              <section className={`inline-flex gap1`}>
                {/*<Typography.Text icon={<IconHome />} link={{ href: '/' }} className={`btn`} />*/}
                <div className={'inline-flex items-center gap2'}>
                  <Link to="/" className={`btn icon-btn`}>
                    {/* <i className="i-ic-baseline-home"></i> */}
                    <AppIcon icon={'i-ic-baseline-home'} />
                  </Link>
                  <Button
                    theme="borderless"
                    onClick={toggleTheme}
                    className={`${classes.btnIcon}`}
                    icon={
                      <AppIcon
                        icon={isDark ? 'i-line-md-moon-rising-filled-loop' : 'i-line-md-sun-rising-filled-loop'}
                      />
                    }
                    style={{ color: 'var(--semi-color-text-2)' }}
                  />
                </div>
                <Divider layout={`vertical`} margin={10} />
                <Dropdown
                  trigger="click"
                  position={`bottomRight`}
                  render={
                    <Dropdown.Menu>
                      {/*<Dropdown.Item*/}
                      {/*  icon={<IconUser />}*/}
                      {/*  onClick={() => nav('/profile')}>*/}
                      {/*  个人中心*/}
                      {/*</Dropdown.Item>*/}
                      {/*<Dropdown.Item*/}
                      {/*  icon={<MaterialSymbolsPasswordRounded />}*/}
                      {/*  onClick={() => setShowUpdatePass(true)}>*/}
                      {/*  修改密码*/}
                      {/*</Dropdown.Item>*/}
                      <Dropdown.Item
                        icon={<MaterialSymbolsLogoutSharp />}
                        onClick={() => logout(() => nav({ to: '/login' }))}>
                        退出
                      </Dropdown.Item>
                    </Dropdown.Menu>
                  }>
                  <Avatar alt="avatar" size={`small`}>
                    {(user.username || 'V')[0].toUpperCase()}
                  </Avatar>
                </Dropdown>
              </section>
            </Nav.Footer>
          </Nav>
        </Header>
        <Layout className={classes.adminMainContent}>
          <Sider className={classes.adminMainSider}>
            <Nav
              className={classes.adminMainSiderNav}
              bodyStyle={{ overflowY: 'auto', height: '100%' }}
              selectedKeys={defaultKeys}
              items={Navs}
              onSelect={(v) => {
                const p = findPathByKey(Navs, v.itemKey as string);
                if (p) nav({ to: p }).then();
              }}
              footer={{
                collapseButton: true,
              }}
            />
          </Sider>
          <Content className={classes.adminMainRightCnt}>
            <Outlet />
          </Content>
        </Layout>
        <Footer className={classes.adminMainFooter}>
          <p>Copyright © {2024}. All Rights Reserved. Version: 0.0.0-dev</p>
        </Footer>
      </Layout>
    </>
  );
};

export const Route = createFileRoute('/admin')({
  component: () => (
    <Suspense fallback={<ContentLoading />}>
      <AdminConfigProvider>
        <AdminLayout />
      </AdminConfigProvider>
    </Suspense>
  ),
  pendingComponent: ContentLoading,
  validateSearch: configSearchSchema,
  beforeLoad: async ({ context, location }) => {
    console.log('🚀 ~ context:', context);
    if (!context.auth || !context.isAdmin) {
      Toast.error('请先登录');
      throw redirect({
        to: '/login',
        search: location.search,
      });
    }
  },
});

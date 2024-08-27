import useTheme from '@/hooks/useTheme.ts';
import Navs, { findKeyByPath, findPathByKey } from '@/pages/navs.tsx';
import { useAuth } from '@/provider/AuthProvider.tsx';
import { Avatar, Button, Divider, Dropdown, Icon, Layout, Nav } from '@douyinfe/semi-ui';
import { createLazyRoute, Outlet, useLocation, useNavigate } from '@tanstack/react-router';
import LineMdMoonRisingFilledLoop from '~icons/line-md/moon-rising-filled-loop';
import LineMdSunRisingFilledLoop from '~icons/line-md/sun-rising-filled-loop';
import MaterialSymbolsLogoutSharp from '~icons/material-symbols/logout-sharp';
import classes from './pages.module.scss';

const { Header, Footer, Sider, Content } = Layout;

const AdminPage = () => {
  console.log('AdminPage render...');
  const { isDark, toggleTheme } = useTheme();
  const loc = useLocation();
  const nav = useNavigate();
  const { user, logout } = useAuth();

  const defaultKeys = useMemo(() => {
    const key = findKeyByPath(Navs, loc.pathname);
    if (key) return [key];
    return ['admin-index'];
  }, [loc]);

  return (
    <Layout className={classes.adminMain}>
      <Header className={classes.adminMainHeader}>
        <Nav mode={`horizontal`}>
          <Nav.Header>
            <img src={`./logo.svg`} alt={`Logo`} style={{ height: '36px', fontSize: 36 }} />
            <span className={'font-bold text-lg ml2'}>V2FD 后台</span>
          </Nav.Header>
          <Nav.Footer>
            <section className={`space-x-2`}>
              {/*<Typography.Text icon={<IconHome />} link={{ href: '/' }} className={`btn`} />*/}
              <a href="/" className={`btn icon-btn`}>
                <i className="i-ic-baseline-home"></i>
              </a>
              {/*<Button*/}
              {/*  theme="borderless"*/}
              {/*  onClick={() => nav({ to: '/' })}*/}
              {/*  className={`${classes.btnIcon}`}*/}
              {/*  icon={<IconHome />}*/}
              {/*  style={{ color: 'var(--semi-color-text-2)' }}*/}
              {/*/>*/}
              <Button
                theme="borderless"
                onClick={toggleTheme}
                className={`${classes.btnIcon}`}
                icon={<Icon svg={isDark ? <LineMdMoonRisingFilledLoop /> : <LineMdSunRisingFilledLoop />} />}
                style={{ color: 'var(--semi-color-text-2)' }}
              />
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
            defaultSelectedKeys={defaultKeys}
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
        <p>Copyright © {2024} V2FD. All Rights Reserved. Version: 0.0.0-dev</p>
      </Footer>
    </Layout>
  );
};

const AdminPageRoute = createLazyRoute('/admin')({
  component: AdminPage,
});
export default AdminPageRoute;

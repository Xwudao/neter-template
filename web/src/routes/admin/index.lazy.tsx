import AdminToolbarTitle from '@/components/admin/layout/AdminToolbarTitle';
import AdminWrapper from '@/components/admin/layout/AdminWrapper';
import AdminSticky from '@/components/admin/layout/AdminSticky';
import { createLazyFileRoute } from '@tanstack/react-router';
import clsx from 'clsx';
import classes from './index.module.scss';

const AdminIndexComponent = () => {
  return (
    <AdminWrapper
      toolbar={
        <div className={classes.toolbarWrapper}>
          <AdminToolbarTitle>后台首页</AdminToolbarTitle>
          <div className={classes.toolbarActions}></div>
        </div>
      }>
      {/* 顶部吸附组件示例 */}
      <AdminSticky position="top" size={'large'} offset={10} className={'p3'}>
        <div className={classes.stickyToolbar}>
          <span className={classes.title}>吸附工具栏</span>
          <div className={classes.actions}>
            <button className={clsx(classes.actionButton, classes.primary)}>刷新</button>
            <button className={clsx(classes.actionButton, classes.success)}>导出</button>
          </div>
        </div>
      </AdminSticky>

      <div className={'px3'}>
        <p>hello index page</p>
        {Array.from({ length: 100 }, (_, i) => (
          <p key={i} className={classes.contentLine}>
            这是第 {i + 1} 行内容 - 当你滚动时，上面的工具栏和下面的状态栏会保持吸附
          </p>
        ))}
      </div>

      {/* 底部吸附组件示例 */}
      {/* <AdminSticky position="bottom">
        <div className={classes.bottomStatus}>
          <span className={classes.statusText}>共 100 条数据 | 页面底部吸附内容</span>
        </div>
      </AdminSticky> */}
    </AdminWrapper>
  );
};

export const Route = createLazyFileRoute('/admin/')({
  component: () => <AdminIndexComponent />,
});

import AdminToolbarTitle from '@/components/admin/layout/AdminToolbarTitle';
import AdminWrapper from '@/components/admin/layout/AdminWrapper';
import ContentLoading from '@/components/loading/ContentLoading';
import { createLazyFileRoute } from '@tanstack/react-router';

const AdminIndexComponent = () => {
  return (
    <AdminWrapper
      toolbar={
        <div className={`line-center justify-between flex-wrap`}>
          <AdminToolbarTitle className={`text-sm`}>后台首页</AdminToolbarTitle>
          <div className={`line-center gap-2 flex-wrap`}></div>
        </div>
      }>
      <p className="">hello index page</p>
      {Array.from({ length: 100 }, (_, i) => (
        <p key={i} className={`text-sm`}>
          这是第 {i + 1} 行内容
        </p>
      ))}
    </AdminWrapper>
  );
};

export const Route = createLazyFileRoute('/admin/')({
  component: () => <AdminIndexComponent />,
  pendingComponent: ContentLoading,
});

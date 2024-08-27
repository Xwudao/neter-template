import AdminContent from '@/pages/admin/layout/AdminContent.tsx';
import AdminToolbarTitle from '@/pages/admin/layout/AdminToolbarTitle.tsx';
import AdminWrapper from '@/pages/admin/layout/AdminWrapper.tsx';
import { adminPageRoute } from '@/router/routes.tsx';
import { createRoute } from '@tanstack/react-router';

type IAdminIndexPage = {};
const AdminIndexPage = (props: IAdminIndexPage) => {
  console.log('AdminIndexPage render...');
  return (
    <>
      <AdminWrapper
        toolbar={
          <div className={`line-center justify-between flex-wrap`}>
            <AdminToolbarTitle className={`text-sm`}>后台首页</AdminToolbarTitle>
            <div className={`line-center gap-2 flex-wrap`}></div>
          </div>
        }>
        <AdminContent>
          <p className="">hello index page</p>
          {/*<div className={`w-100`}>*/}
          {/*  <UploadCard onSuccess={(rtn) => console.log(rtn)} />*/}
          {/*</div>*/}
        </AdminContent>
      </AdminWrapper>
    </>
  );
};

const AdminIndexPageRoute = createRoute({
  getParentRoute: () => adminPageRoute,
  component: () => <AdminIndexPage />,

  path: '/',
});

export default AdminIndexPageRoute;

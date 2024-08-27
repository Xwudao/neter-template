import { FC, PropsWithChildren } from 'react';
import classes from './pages.module.scss';


type IAdminToolbar = {};
const AdminToolbar: FC<PropsWithChildren<IAdminToolbar>> = ({children}) => {
 console.log('adminToolbar render...');
 return (
  <div className={classes.adminToolbar}>
    {children}
  </div>
 );
};

export default AdminToolbar;

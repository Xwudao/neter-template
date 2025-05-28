import { FC, PropsWithChildren } from 'react';
import classes from './styles.module.scss';

type IAdminToolbarTitle = {
  className?: string;
};
const AdminToolbarTitle: FC<PropsWithChildren<IAdminToolbarTitle>> = ({ children, className = '' }) => {
  console.log('adminToolbarTitle render...');
  return <h2 className={`${classes.adminToolbarTitle} ${className}`}>{children}</h2>;
};

export default AdminToolbarTitle;

import { Spin } from '@douyinfe/semi-ui';
import clsx from 'clsx';

import classes from './loading.module.scss';

interface Props {
  text?: string;
  className?: string;
}

function PageLoading(props: Props) {
  const { text = 'Loading...', className = '' } = props;

  return (
    <div className={clsx('w-screen h-screen flex items-center justify-center gap1 text-primary', classes.pageLoading)}>
      <Spin size="large" />
      <span className={clsx(className)}>{text}</span>
    </div>
  );
}

export default PageLoading;

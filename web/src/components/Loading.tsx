import classnames from 'classnames';
import { FC, PropsWithChildren } from 'react';

type ILoading = {
  className?: string;
  text?: string;
  show?: boolean;
};
const Loading: FC<PropsWithChildren<ILoading>> = ({ show = true, className = '', text = '加载中...' }) => {
  console.log('loading render...');
  if (!show) return null;
  return (
    <div className={classnames('text-pending text-base', className)}>
      <i className="i-svg-spinners-bars-rotate-fade"></i>
      <span className={`am ml1 select-none`}>{text}</span>
    </div>
  );
};

export default Loading;

import AppIcon from '@/components/AppIcon';
import clsx from 'clsx';

interface Props {
  text?: string;
  className?: string;
}

function PageLoading(props: Props) {
  const { text = 'Loading...', className = '' } = props;

  return (
    <div className={clsx('w-screen h-screen flex items-center justify-center gap1 text-primary')}>
      <AppIcon icon={'i-svg-spinners-clock'} />
      <span className={clsx(className)}>{text}</span>
    </div>
  );
}

export default PageLoading;

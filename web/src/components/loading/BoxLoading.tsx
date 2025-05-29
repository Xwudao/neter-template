import AppIcon from '@/components/AppIcon';
import clsx from 'clsx';

interface Props {
  text?: string;
  className?: string;
  size?: 'sm' | 'md' | 'lg';
}

function BoxLoading(props: Props) {
  const { text = 'Loading...', className = '', size = 'md' } = props;

  const sizeClasses = {
    sm: 'p-2',
    md: 'p-4',
    lg: 'p-6',
  };

  return (
    <div className={clsx('flex flex-col items-center justify-center gap-2 text-primary', sizeClasses[size], className)}>
      <AppIcon icon={'i-svg-spinners-clock'} className="" />
      <span className="text-sm">{text}</span>
    </div>
  );
}

export default BoxLoading;

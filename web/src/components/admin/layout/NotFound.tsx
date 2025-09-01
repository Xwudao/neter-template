import { Button, Empty } from '@douyinfe/semi-ui';
import { useNavigate } from '@tanstack/react-router';
import classes from './styles.module.scss';

import useTheme from '@/hooks/useTheme';
import { IllustrationNotFound, IllustrationNotFoundDark } from '@douyinfe/semi-illustrations';

function NotFound() {
  // const {} = props

  const nav = useNavigate();
  const { isDark } = useTheme();
  const Svg = isDark ? IllustrationNotFoundDark : IllustrationNotFound;

  return (
    <div className={classes['not-found-container']}>
      <div className={classes['not-found-content']}>
        <Empty image={<Svg />} className={classes['not-found-image']} />
        <h2 className={classes['not-found-title']}>你来到了一片荒芜</h2>
        <p className={classes['not-found-description']}>抱歉，您访问的页面不存在</p>
        <div className={classes['not-found-actions']}>
          <Button type={'secondary'} onClick={() => (window.location.href = '/')}>
            返回首页
          </Button>
          <Button onClick={() => nav({ to: '..' })}>返回上页</Button>
        </div>
      </div>
    </div>
  );
}

export default NotFound;

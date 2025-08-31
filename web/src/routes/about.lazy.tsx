import { createLazyFileRoute } from '@tanstack/react-router';
import classes from './about.module.scss';

const AboutPage = () => {
  return (
    <div className={classes.container}>
      <div className={classes.mainCard}>
        <h1 className={classes.title}>关于我们</h1>
        <p className={classes.description}>这是关于页面的简单介绍。我们致力于为用户提供优质的产品和服务。</p>
      </div>
    </div>
  );
};

export const Route = createLazyFileRoute('/about')({
  component: AboutPage,
});

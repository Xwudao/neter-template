import useTheme from '@/hooks/useTheme';
import useConfig from '@/provider/useConfig';
import { Button, Input } from '@douyinfe/semi-ui';
import { createLazyFileRoute, useNavigate } from '@tanstack/react-router';
import styles from './index.module.scss';

export const Route = createLazyFileRoute('/')({
  component: () => <IndexComponent />,
});

const IndexComponent = () => {
  const nav = useNavigate();
  const { toggleTheme } = useTheme();
  const {
    config: { site_info },
  } = useConfig();

  return (
    <div className={styles.container}>
      <div className={styles.mainCard}>
        <h1 className={styles.title}>Welcome to {site_info.site_name}</h1>
        <p className={styles.description}>
          This is the neter template starter page. Explore the features and navigate through different sections.
        </p>

        <div className={styles.actions}>
          <div className={styles.themeToggle}>
            <Button onClick={toggleTheme} theme="borderless" size="large">
              🌓 Toggle Theme
            </Button>
          </div>

          <hr className={styles.divider} />

          <div className={styles.inputSection}>
            <span className={styles.inputLabel}>Try the input component:</span>
            <Input className={styles.demoInput} placeholder="Enter something here..." size="large" />
          </div>

          <hr className={styles.divider} />

          <div className={styles.buttonGroup}>
            <Button className={styles.navButton} type="primary" size="large" onClick={() => nav({ to: '/about' })}>
              📖 About
            </Button>
            <Button className={styles.navButton} type="secondary" size="large" onClick={() => nav({ to: '/login' })}>
              🔐 Login
            </Button>
          </div>

          <div className={styles.siteInfo}>
            <div className={styles.siteName}>{site_info.site_name}</div>
            <div className={styles.siteKeywords}>{site_info.site_keywords}</div>
          </div>
        </div>
      </div>
    </div>
  );
};

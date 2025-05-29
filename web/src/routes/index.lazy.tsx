import useTheme from '@/hooks/useTheme';
import useConfig from '@/provider/useConfig';
import { Button, Input } from '@douyinfe/semi-ui';
import { createLazyFileRoute, useNavigate } from '@tanstack/react-router';
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
    <div>
      <h1>Index Page</h1>
      <p>This is the index page of the application.</p>

      <Button onClick={toggleTheme}>
        <span>Toggle Theme</span>
      </Button>
      <hr />

      <Input />

      <button
        onClick={() => {
          nav({ to: '/about' });
        }}>
        to about
      </button>
      <button
        onClick={() => {
          nav({ to: '/login' });
        }}>
        login
      </button>

      <hr />

      {site_info.site_name} <br/>
      {site_info.site_keywords}
    </div>
  );
};

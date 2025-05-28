import useTheme from '@/hooks/useTheme';
import { Button, Input } from '@douyinfe/semi-ui';
import { createLazyFileRoute, useNavigate } from '@tanstack/react-router';
export const Route = createLazyFileRoute('/')({
  component: () => <IndexComponent />,
});

const IndexComponent = () => {
  const nav = useNavigate();
  const { toggleTheme } = useTheme();
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
      {/* <Chart /> */}
    </div>
  );
};

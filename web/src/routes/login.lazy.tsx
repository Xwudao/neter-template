import { postApiUserLogin } from '@/api/userApi';
import { onError, onSuccess } from '@/core/callback';
import { UserRole } from '@/core/constants';
import useAuth from '@/provider/useAuth';
import { Button, Form, Space } from '@douyinfe/semi-ui';
import { useMutation } from '@tanstack/react-query';
import { createLazyFileRoute, Link, useNavigate } from '@tanstack/react-router';
import clsx from 'clsx';
import classes from './login.module.scss';
const { Input } = Form;
const Login = () => {
  const { mutate, isPending } = useMutation({
    mutationKey: ['login'],
    mutationFn: postApiUserLogin,
  });

  const { toLogin, logged, isAdmin } = useAuth();
  const nav = useNavigate();

  useEffect(() => {
    if (!logged) return;

    if (isAdmin) {
      nav({ to: '/admin', search: { tab: '' } }).then(() => {});
    } else {
      nav({ to: '/' });
    }
  }, [logged, nav, isAdmin]);

  const handleLogin = (values: any) => {
    mutate(values, {
      onSuccess: onSuccess('登录成功', (rtn) => {
        toLogin(
          {
            ...rtn.user,
            token: rtn.token,
          },
          () => {
            setTimeout(() => {
              if (rtn.user.role?.includes(UserRole.ADMIN)) {
                nav({ to: '/admin', search: { tab: '' } });
              } else {
                nav({ to: '/' });
              }
            }, 0);
          },
        );
      }),
      onError: onError(),
    });
  };

  return (
    <section className={clsx(classes.loginContainer)}>
      <div className={clsx(classes.loginBox)}>
        <h1>登录系统</h1>
        <Form layout={'vertical'} onSubmit={handleLogin}>
          <Input
            field="username"
            label="用户名"
            placeholder="请输入用户名"
            rules={[{ required: true, message: '请输入用户名' }]}
          />
          <Input
            field="password"
            label="密码"
            mode="password"
            rules={[{ required: true, message: '请输入密码' }]}
            placeholder="请输入密码"
          />
          <Button htmlType={'submit'} type="primary" block loading={isPending}>
            登录
          </Button>
        </Form>

        <div className={clsx(classes.navigationButtons)}>
          <Space spacing={'loose'}>
            <Link to="/">返回首页</Link>
            <Link to="/register">注册账号</Link>
            <Link to="..">返回上页</Link>
          </Space>
        </div>
      </div>
    </section>
  );
};

export const Route = createLazyFileRoute('/login')({
  component: () => <Login />,
});

import { Button, Form, Space } from '@douyinfe/semi-ui';
import { IconHome, IconChevronLeft } from '@douyinfe/semi-icons';
import { createLazyFileRoute, useNavigate, useRouter } from '@tanstack/react-router';
import clsx from 'clsx';
import classes from './login.module.scss';
import { useMutation } from '@tanstack/react-query';
import { postApiUserLogin } from '@/api/userApi';
import { onError, onSuccess } from '@/core/callback';
import useAuth from '@/provider/useAuth';
import { UserRole } from '@/core/constants';
const { Input } = Form;
const Login = () => {
  const { mutate, isPending } = useMutation({
    mutationKey: ['login'],
    mutationFn: postApiUserLogin,
  });

  const { toLogin } = useAuth();
  const nav = useNavigate();
  const router = useRouter();
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

  const handleGoHome = () => {
    nav({ to: '/' });
  };

  const handleGoBack = () => {
    router.history.back();
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
            size="large"
            rules={[{ required: true, message: '请输入用户名' }]}
          />
          <Input
            field="password"
            label="密码"
            mode="password"
            size="large"
            rules={[{ required: true, message: '请输入密码' }]}
            placeholder="请输入密码"
          />
          <Button htmlType={'submit'} type="primary" size="large" block loading={isPending}>
            登录
          </Button>
        </Form>

        <div className={clsx(classes.navigationButtons)}>
          <Space>
            <Button icon={<IconChevronLeft />} onClick={handleGoBack} type="tertiary">
              返回上页
            </Button>
            <Button icon={<IconHome />} onClick={handleGoHome} type="tertiary">
              返回首页
            </Button>
          </Space>
        </div>
      </div>
    </section>
  );
};

export const Route = createLazyFileRoute('/login')({
  component: () => <Login />,
});

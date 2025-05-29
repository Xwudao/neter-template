import { Button, Form } from '@douyinfe/semi-ui';
import { createLazyFileRoute, useNavigate } from '@tanstack/react-router';
import clsx from 'clsx';
import classes from './styles.module.scss';
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
    <section className={clsx('w-screen h-screen flex items-center justify-center')}>
      <div className={clsx(classes.loginBox)}>
        <h1 className={'text-lg fw-bold mb3'}>登录系统</h1>
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
          <Button htmlType={'submit'} type="primary" className="mt-4" block loading={isPending}>
            登录
          </Button>
        </Form>
      </div>
    </section>
  );
};

export const Route = createLazyFileRoute('/login')({
  component: () => <Login />,
});

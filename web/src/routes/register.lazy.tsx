import { Button, Form, Space } from '@douyinfe/semi-ui';
import { createLazyFileRoute, Link } from '@tanstack/react-router';
import clsx from 'clsx';
import classes from './register.module.scss';

const { Input } = Form;

const Register = () => {
  const handleRegister = (values: any) => {
    console.log('Register values:', values);
    // TODO: Implement register API call when available
  };

  return (
    <section className={clsx(classes.registerContainer)}>
      <div className={clsx(classes.registerBox)}>
        <h1>用户注册</h1>
        <Form layout={'vertical'} onSubmit={handleRegister}>
          <Input
            field="username"
            label="用户名"
            placeholder="请输入用户名"
            rules={[
              { required: true, message: '请输入用户名' },
              { min: 3, message: '用户名至少3个字符' },
            ]}
          />
          <Input
            field="email"
            label="邮箱"
            placeholder="请输入邮箱"
            rules={[
              { required: true, message: '请输入邮箱' },
              { type: 'email', message: '请输入有效的邮箱地址' },
            ]}
          />
          <Input
            field="password"
            label="密码"
            mode="password"
            placeholder="请输入密码"
            rules={[
              { required: true, message: '请输入密码' },
              { min: 6, message: '密码至少6个字符' },
            ]}
          />
          <Input
            field="confirmPassword"
            label="确认密码"
            mode="password"
            placeholder="请再次输入密码"
            rules={[
              { required: true, message: '请确认密码' },
              {
                validator: (rule, value, callback, source, options) => {
                  if (value && value !== source.password) {
                    return new Error('两次输入的密码不一致');
                  }
                  return true;
                },
              },
            ]}
          />
          <Button htmlType={'submit'} type="primary" block>
            注册
          </Button>
        </Form>

        <div className={clsx(classes.navigationButtons)}>
          <Space spacing={'loose'}>
            <div>
              已有账号？
              <Link to="/login">立即登录</Link>
            </div>
            <Link to="/">返回首页</Link>
          </Space>
        </div>
      </div>
    </section>
  );
};

export const Route = createLazyFileRoute('/register')({
  component: () => <Register />,
});

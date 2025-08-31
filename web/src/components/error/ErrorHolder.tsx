import { ErrorComponentProps } from '@tanstack/react-router';
import React from 'react';
import classes from './error-holder.module.scss';

interface ErrorHolderProps extends ErrorComponentProps {
  title?: string;
  showRetry?: boolean;
  onRetry?: () => void;
}

export const ErrorHolder: React.FC<ErrorHolderProps> = ({ error, info, reset, title, showRetry = true, onRetry }) => {
  const handleRetry = () => {
    if (onRetry) {
      onRetry();
    } else if (reset) {
      reset();
    }
  };

  const getErrorMessage = () => {
    if (error?.message) return error.message;
    if (typeof error === 'string') return error;
    return '发生了未知错误';
  };

  const getErrorTitle = () => {
    if (title) return title;
    if (error?.message?.includes('404') || error?.message?.includes('Not Found')) {
      return '页面未找到';
    }
    return '出错了';
  };

  return (
    <div className={classes.error_holder}>
      <div className={classes.error_content}>
        <div className={classes.error_icon}>⚠️</div>
        <h3 className={classes.error_title}>{getErrorTitle()}</h3>
        <p className={classes.error_message}>{getErrorMessage()}</p>

        {showRetry && (
          <button className={classes.retry_button} onClick={handleRetry}>
            重试
          </button>
        )}

        {process.env.NODE_ENV === 'development' && info && (
          <details className={classes.error_details}>
            <summary>错误详情</summary>
            <pre className={classes.error_stack}>{info.componentStack || error?.stack}</pre>
          </details>
        )}
      </div>
    </div>
  );
};

// Export as default for TanStack Router
export const DefaultErrorComponent = (props: ErrorHolderProps) => {
  return <ErrorHolder {...props} />;
};

export default DefaultErrorComponent;

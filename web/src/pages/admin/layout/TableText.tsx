import { Tooltip, Typography } from '@douyinfe/semi-ui';
import { RenderContent } from '@douyinfe/semi-ui/lib/es/tooltip';
import { TypographyBaseSize } from '@douyinfe/semi-ui/lib/es/typography/interface';
import { FC, PropsWithChildren, ReactNode, useCallback } from 'react';
type ITableText = {
  text?: string;
  width?: number;
  size?: TypographyBaseSize;
};
const TableText: FC<PropsWithChildren<ITableText>> = ({ text, size = `small`, width = 320 }) => {
  console.log('tableText render...');
  const customRenderTooltip = useCallback((content: ReactNode | RenderContent, children: ReactNode) => {
    return <Tooltip content={content}>{children}</Tooltip>;
  }, []);
  return (
    <Typography.Text
      size={size}
      style={{ width }}
      ellipsis={{
        showTooltip: {
          renderTooltip: customRenderTooltip,
        },
      }}>
      {text}
    </Typography.Text>
  );
};

export default TableText;

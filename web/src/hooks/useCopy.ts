import { Toast } from '@douyinfe/semi-ui';
import { useCopyToClipboard } from 'react-use';
import { useEffect } from 'react';

const useCopy = () => {
  const [state, copy] = useCopyToClipboard();
  useEffect(() => {
    if (state.value) {
      if (state.value.length < 5) {
        Toast.success(`复制成功: ${state.value}`);
      } else {
        Toast.success(`复制成功`);
      }

      return;
    }
    if (state.error) {
      Toast.error(String(state.error));
    }
  }, [state]);

  return copy;
};

export default useCopy;

import { Category } from '@/api/categoryApi.ts';
import { postApiCreateSite } from '@/api/siteApi.ts';
import { onError, onSuccess } from '@/core/callback.ts';
import useCopy from '@/hooks/useCopy.ts';
import FormWrapper from '@/pages/admin/popovers/FormWrapper.tsx';
import SiteForm from '@/pages/admin/popovers/SiteForm.tsx';
import { extractUrls } from '@/utils/url.ts';
import { Button, Divider, Select, TextArea } from '@douyinfe/semi-ui';
import { useMutation } from '@tanstack/react-query';
import { FC, PropsWithChildren } from 'react';

type IAddBulkSite = {
  categories: Category[];
  onAdded?: () => void;
};
const AddBulkSite: FC<PropsWithChildren<IAddBulkSite>> = ({ categories, onAdded }) => {
  console.log('addMutiSite render...');
  const [text, setText] = useState('');
  const [nowUrl, setNowUrl] = useState<string>();
  const [urls, setUrls] = useState([] as string[]);

  const copy = useCopy();

  const { mutate: doAdd, isPending } = useMutation({
    mutationFn: postApiCreateSite,
  });
  const handleAdd = (values: any) => {
    doAdd(
      {
        ...values,
        keywords: values.keywords.split(','),
      },
      {
        onSuccess: onSuccess(true, () => onAdded?.()),
        onError: onError(),
      },
    );
  };

  return (
    <div className={`p2 space-y-3 w-130`}>
      <h2 className={`text-base fw-bold`}>批量解析站点</h2>
      <div className={`w-100 space-y-3`}>
        <TextArea
          value={text}
          onBlur={() => {
            const urls = extractUrls(
              text,
              (u) => !u.match(/(jpg|png|jpeg|gif|svg|html)$/),
            );
            setUrls(urls);
          }}
          onChange={(value) => setText(value)}
          placeholder={`输入html代码段`}
        />
        <Select
          value={nowUrl}
          onChange={(value) => {
            copy(value as string);
            setNowUrl(value as string);
          }}
          className={`w-full max-w-full`}
          position={`bottomRight`}
          optionList={urls.map((c) => ({ label: c, value: c }))}
          placeholder={`解析的地址`}
        />
      </div>
      <Divider margin={10} />
      <div>
        <FormWrapper
          className={`grid grid-cols-2 gap-3`}
          title={`添加站点`}
          onSubmit={handleAdd}>
          <SiteForm categories={categories} />
          <Button block loading={isPending} htmlType={`submit`}>
            提交
          </Button>
        </FormWrapper>
      </div>
    </div>
  );
};

export default AddBulkSite;

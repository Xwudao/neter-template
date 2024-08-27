import { Category } from '@/api/categoryApi.ts';
import { postApiCreateSite } from '@/api/siteApi.ts';
import { onError, onSuccess } from '@/core/callback.ts';
import FormWrapper from '@/pages/admin/popovers/FormWrapper.tsx';
import SiteForm from '@/pages/admin/popovers/SiteForm.tsx';
import { Button } from '@douyinfe/semi-ui';
import { useMutation } from '@tanstack/react-query';
import { isArray } from 'radash';
import { FC, PropsWithChildren } from 'react';

type IAddSite = {
  categories: Category[];
  onAdded?: () => void;
};
const AddSite: FC<PropsWithChildren<IAddSite>> = ({ categories, onAdded }) => {
  console.log('addSite render...');
  const { mutate: doAdd, isPending } = useMutation({
    mutationFn: postApiCreateSite,
  });
  const handleAdd = (values: any) => {
    doAdd(
      {
        ...values,
        keywords: isArray(values.keywords) ? values.keywords : values.keywords.split(','),
      },
      {
        onSuccess: onSuccess(true, () => onAdded?.()),
        onError: onError(),
      },
    );
  };
  return (
    <div className={`p2 w-130`}>
      <FormWrapper className={`grid grid-cols-2 gap-3`} title={`添加站点`} onSubmit={handleAdd}>
        <SiteForm categories={categories} />
        <Button block loading={isPending} htmlType={`submit`}>
          提交
        </Button>
      </FormWrapper>
    </div>
  );
};

export default AddSite;

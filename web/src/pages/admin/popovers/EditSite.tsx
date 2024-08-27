import { Category } from '@/api/categoryApi.ts';
import { postApiUpdateSite, Site } from '@/api/siteApi.ts';
import { onError, onSuccess } from '@/core/callback.ts';
import FormWrapper from '@/pages/admin/popovers/FormWrapper.tsx';
import SiteForm from '@/pages/admin/popovers/SiteForm.tsx';
import { Button } from '@douyinfe/semi-ui';
import { useMutation } from '@tanstack/react-query';
import { isArray } from 'radash';
import { FC, PropsWithChildren } from 'react';

type IEditSite = {
  data: Site;
  categories: Category[];
  onEdited?: () => void;
};
const EditSite: FC<PropsWithChildren<IEditSite>> = ({ data, categories, onEdited }) => {
  console.log('editSite render...', data);
  const { mutate: doEdit, isPending } = useMutation({
    mutationFn: postApiUpdateSite,
  });
  const handleEdit = (values: any) => {
    doEdit(
      { ...values, keywords: isArray(values.keywords) ? values.keywords : values.keywords.split(',') },
      {
        onSuccess: onSuccess(true, () => onEdited?.()),
        onError: onError(),
      },
    );
  };
  return (
    <div className={`p2 w-120`}>
      <FormWrapper
        className={`grid grid-cols-2 gap-3`}
        title={`修改站点`}
        onSubmit={handleEdit}
        initValues={{
          ...data,
          category_id: data.edges?.cates.map((c) => c.id),
        }}>
        <SiteForm categories={categories} />
        <Button block loading={isPending} htmlType={`submit`}>
          提交
        </Button>
      </FormWrapper>
    </div>
  );
};

export default EditSite;

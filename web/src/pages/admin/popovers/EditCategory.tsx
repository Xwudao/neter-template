import { Category, postApiUpdateCategory } from '@/api/categoryApi.ts';
import { onError, onSuccess } from '@/core/callback.ts';
import CategoryForm from '@/pages/admin/popovers/CategoryForm.tsx';
import FormWrapper from '@/pages/admin/popovers/FormWrapper.tsx';
import { Button } from '@douyinfe/semi-ui';
import { useMutation } from '@tanstack/react-query';
import { FC, PropsWithChildren } from 'react';

type IEditCategory = {
  category: Category;
  categories: Category[];
  onEdited?: () => void;
};
const EditCategory: FC<PropsWithChildren<IEditCategory>> = ({ categories, category, onEdited }) => {
  console.log('editCategory render...');
  const { mutate: doEdit, isPending } = useMutation({
    mutationFn: postApiUpdateCategory,
  });
  const handleEdit = (values: any) => {
    doEdit(values, {
      onSuccess: onSuccess(true, () => onEdited?.()),
      onError: onError(),
    });
  };
  return (
    <div className={`p2 w-62`}>
      <FormWrapper title={`修改分类`} onSubmit={handleEdit} initValues={category}>
        <CategoryForm categories={categories} category={category} />
        <Button block loading={isPending} htmlType={`submit`}>
          提交
        </Button>
      </FormWrapper>
    </div>
  );
};

export default EditCategory;

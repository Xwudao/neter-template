import { Category } from '@/api/categoryApi.ts';
import { getApiGetSiteMeta, postApiUploadIcon } from '@/api/siteApi.ts';
import { onSuccess } from '@/core/callback.ts';
import { KEY_CLEAN_URL_AUTO } from '@/core/constants.ts';
import { ICON_PIPE } from '@/core/types.ts';
import useCopy from '@/hooks/useCopy.ts';
import { cleanUrl, getHostname } from '@/utils/path.ts';
import { arrayToTree } from '@/utils/tree.ts';
import { IconRefresh, IconUpload } from '@douyinfe/semi-icons';
import { Button, Form, Typography, useFormApi, useFormState } from '@douyinfe/semi-ui';
import { useMutation } from '@tanstack/react-query';
import { FC, PropsWithChildren } from 'react';
import { useLocalStorage } from 'react-use';

const { TreeSelect, TextArea, Slider, Checkbox, Input } = Form;

type ISiteForm = {
  categories: Category[];
};
const SiteForm: FC<PropsWithChildren<ISiteForm>> = ({ categories }) => {
  console.log('siteForm render...');

  const [autoClean, setAutoClean] = useLocalStorage(KEY_CLEAN_URL_AUTO, false);
  const copy = useCopy();

  const formApi = useFormApi();
  const formState = useFormState();

  const treeData = useMemo(() => {
    const data = categories.map((item) => ({
      parent_id: item.parent_id,
      label: item.name,
      value: item.id,
      id: item.id,
      key: `${item.parent_id || 0}-${item.id}`,
    }));
    return arrayToTree(data);
  }, [categories]);

  const iconUrl = useMemo(() => {
    if (formState.values['icon']?.startsWith('http')) return formState.values['icon'];
    if (formState.values['icon']?.startsWith('favicon/')) return ICON_PIPE + formState.values['icon'];
    return '/static/logo.svg';
  }, [formState]);

  const { mutate: doGetMeta, isPending } = useMutation({
    mutationFn: getApiGetSiteMeta,
  });

  const handleRefresh = () => {
    doGetMeta(
      {
        url: formApi.getValue(`url`),
      },
      {
        onSuccess: onSuccess('采集成功', (data) => {
          if (data.title) formApi.setValue('name', data.title);
          if (data.keywords) formApi.setValue('keywords', data.keywords);
          if (data.description) formApi.setValue('description', data.description);
          if (data.favicon) formApi.setValue('icon', data.favicon);
        }),
      },
    );
  };
  const { mutate: doUpload, isPending: uploading } = useMutation({
    mutationFn: postApiUploadIcon,
  });
  const handleUpload = () => {
    doUpload(
      {
        icon_url: formApi.getValue(`icon`),
        site_url: formApi.getValue(`url`),
      },
      {
        onSuccess: onSuccess('上传成功', (data) => {
          formApi.setValue('icon', data);
        }),
      },
    );
  };

  return (
    <>
      <Input
        field={`url`}
        label={`站点地址`}
        placeholder={`请输入站点地址`}
        rules={[{ required: true, message: '必填' }]}
        onBlur={() => {
          if (autoClean) {
            formApi.setValue('url', cleanUrl(formState.values['url']));
          }
        }}
        extraText={
          <p>
            <Typography.Text
              size={`small`}
              link
              onClick={() => {
                formApi.setValue('url', cleanUrl(formState.values['url']));
              }}>
              清理URL
            </Typography.Text>
          </p>
        }
        addonAfter={
          <Button
            disabled={!formState.values['url']}
            loading={isPending}
            onClick={handleRefresh}
            icon={<IconRefresh />}
          />
        }
      />
      <Input field={`name`} label={`名称`} placeholder={`请输入名称`} rules={[{ required: true, message: '必填' }]} />
      <Input
        field={`keywords`}
        label={`关键词`}
        placeholder={`请输入关键词,分隔`}
        rules={[{ required: true, message: '必填' }]}
      />
      <Input
        field={`icon`}
        label={`图标`}
        placeholder={`请输入图标`}
        rules={[{ required: true, message: '必填' }]}
        extraText={
          <p className={`flex gap-2`}>
            <Typography.Text
              size={`small`}
              disabled={!formState.values['url']}
              link
              onClick={() => {
                formApi.setValue('icon', `https://api.iowen.cn/favicon/${getHostname(formState.values['url'])}.png`);
              }}>
              Api①
            </Typography.Text>
            <Typography.Text
              size={`small`}
              disabled={!formState.values['url']}
              link
              onClick={() => {
                formApi.setValue(
                  'icon',
                  `https://www.google.com/s2/favicons?domain=${getHostname(formState.values['url'])}&sz=128`,
                );
              }}>
              Api②
            </Typography.Text>
          </p>
        }
        addonBefore={
          <img width={30} height={30} className={`px1 cp`} onClick={() => copy(iconUrl)} src={iconUrl} alt={`icon`} />
        }
        addonAfter={
          <Button
            disabled={!formState.values['icon']}
            onClick={handleUpload}
            loading={uploading}
            icon={<IconUpload />}
          />
        }
      />
      <TreeSelect
        className={`w-full`}
        placeholder={`请选择分类`}
        field={`category_id`}
        label={`请选择分类`}
        defaultExpandAll
        maxTagCount={1}
        checkRelation="unRelated"
        multiple
        rules={[{ required: true, message: '必填' }]}
        treeData={treeData}
      />
      <TextArea
        field={`description`}
        rules={[{ required: true, message: '必填' }]}
        label={`描述`}
        placeholder={`请输入描述`}
      />
      <Slider label={`权重排序`} field={`item_order`} min={1} max={100} showBoundary={false} />
      <Checkbox noLabel field={`clean`} initValue={autoClean} onChange={(v) => setAutoClean(v.target.checked)}>
        自动清理URL
      </Checkbox>
    </>
  );
};

export default SiteForm;

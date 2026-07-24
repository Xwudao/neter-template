# 前端基础设施迁移 — 优化提示词模板

> 适用于：将现有项目的前端基础设施迁移到新项目，或在新项目中搭建完整的前端体系。

---

## 核心原则

```
1. 先分析，后实施。先完整分析参考项目，整理出值得迁移的清单。
2. 不机械复制。每条代码都要审查、适配目标项目。
3. 保持目标项目现有风格和目录结构。
4. 只引入真正会用到的依赖。
5. 每完成一个阶段都检查能否正常工作。
6. 可优化的地方在保持兼容的前提下改进。
```

---

## 一、分析阶段 Prompt

```
目标：完整分析 [参考项目]，整理可迁移的前端基础设施。

要求：
- 不要机械复制代码
- 优先复用参考项目已有实现
- 保持 [目标项目] 当前目录结构和代码风格
- 所有迁移内容必须适配当前项目
- 不引入无用依赖

请先完整分析 [参考项目]，整理以下清单：

### 技术栈清单
- 构建工具（Vite / Webpack / 其他）
- CSS 方案（UnoCSS / Tailwind / CSS Modules / SCSS）
- UI 框架（React / Vue / 其他）
- Router
- 状态管理
- 类型检查
- Linter / Formatter
- 测试框架

### 基础设施清单
- UnoCSS 配置（presets / shortcuts / rules / theme / safelist / transformers）
- Icons 方案（unplugin-icons / icon collections / auto install）
- Auto Import（unplugin-auto-import / resolver / d.ts）
- Router 配置（route tree / error boundary / not found / pending / loader / devtools）
- 状态管理组织方式（store / slice / persist）
- Provider 层次结构
- 样式系统（design tokens / theme / dark mode / breakpoints）

### 确认为每个项目分别列出 ✅值得迁移 / ❌不迁移 / ⚠️有条件迁移
```

---

## 二、依赖安装 Prompt

```
目标：安装所有确认迁移的依赖。

要求：
- 只安装真正会用到的包
- 区分 dependencies 和 devDependencies
- 版本号与参考项目保持一致或更高

运行 pnpm add ...（dependencies）
运行 pnpm add -D ...（devDependencies）
```

---

## 三、配置文件迁移 Prompt

```
目标：迁移 [参考项目] 的配置文件到 [目标项目]。

要求：
- 每条配置都要理解用途
- 删除目标项目不需要的部分
- 保持目标项目现有的其他配置完整

配置文件清单：
- vite.config.ts（UnoCSS / AutoImport / Icons / Router 插件、resolve alias、SCSS additionalData、css modules 配置）
- tsconfig.json / tsconfig.app.json（paths alias、lib、types）
- uno.config.ts（presets、shortcuts、rules、safelist）
- .oxlintrc.json（plugins、rules、ignorePatterns）
- .oxfmtrc.json（semi、singleQuote、trailingComma、sortImports、ignorePatterns）
- .vscode/settings.json（formatter、oxc 配置）
- .env（环境变量）
- package.json（scripts）

注意：
- ❌ 不要复制构建无关的插件（如混淆、压缩、自定义插件等）
- ❌ 不要复制无法确定用途的配置
- ❌ 不要复制与项目无关的 build.rollupOptions / build.target
- ✅ 确保 SCSS `additionalData` 中的 mixin 路径正确（如 `@/styles/mixins`）
- ✅ 确保 UnoCSS configFile 路径为绝对路径
```

---

## 四、样式系统 Prompt

```
目标：建立完整的 SCSS Design Token 系统。

要求：
- 全部基于 CSS Variables，禁止写死颜色
- Dark 为默认主题，Light 通过 `[data-theme='light']` 覆盖
- 与 ThemeProvider / Zustand 联动

必须定义：

### Color — 完整颜色体系
- Surface: canvas / surface / surface-muted / surface-elevated / input
- Text: text / text-muted / text-strong
- Border: border-soft / border-strong
- Accent: accent / accent-hover / accent-active / accent-soft / accent-border / accent-contrast
- Semantic: success / warning / danger / info / signal
- 所有 semantic 颜色附带 -soft 变体
- Hover overlays: hover-overlay / active-overlay
- Decoration: grid-line / backdrop-glow

### Typography
- font-sans / font-heading / font-mono
- font-weight-regular / medium / semibold / bold
- font-size-2xs / xs / sm / md / lg / xl / 2xl / 3xl / 4xl / display
- line-height-tight / snug / normal / relaxed
- letter-spacing-tight / normal / wide

### Spacing (Tailwind 风格)
- space-0 / 0-5 / 1 / 1-5 / 2 / 2-5 / 3 / 4 / 5 / 6 / 7 / 8 / 9 / 10 / 12 / 16
- padding-control-xs / sm / md / lg
- padding-card / padding-page / padding-section
- stack-gap-xs / sm / md / lg

### Radius
- radius-none / xs / sm / md / lg / xl / 2xl / pill

### Shadow
- shadow-xs / sm / md / lg

### ZIndex
- z-raised / dropdown / header / modal / top / toast

### Motion
- duration-fast / normal / slow
- ease-standard

### Layout
- container-page / header-h / sidebar-w

### SCSS Mixins（在 _mixins.scss 中）
- Token accessors: color-token / space / radius / font-size
- focus-ring / interactive-surface / card-surface / elevated-surface
- text-truncate / text-clamp / flex-center / flex-between
- 响应式: sm / md / lg / xl
```

---

## 五、Store + Provider Prompt

```
目标：迁移状态管理和 Provider 层。

要求：
- Zustand store 与 persist 中间件联动
- ThemeProvider 负责：
  1. 监听 system color scheme 变化
  2. data-theme / data-theme-mode / color-scheme 同步到 document
  3. Accent 颜色写入 CSS Variables
  4. 提供 Toaster（react-hot-toast）
- 导出 useTheme() hook
- Provider 放在 main.tsx 中包裹 RouterProvider

参考结构：
```

src/
store/useAppConfig.ts # Zustand + persist
provider/ThemeProvider.tsx # Context + DOM 同步
main.tsx # 入口

```

注意：
- partialize 只持久化需要的字段
- merge / migrate 处理版本升级
- Accent 预设保持 7 种：indigo / emerald / amber / rose / sky / violet / orange
```

---

## 六、Router Prompt

```
目标：迁移 TanStack Router。

配置：
- router.tsx：createRouter + createBrowserHistory
- routes/__root.tsx：根路由 + Outlet + Devtools（仅开发环境）
- routes/index.tsx：首页路由
- routeTree.gen.ts：手动或自动生成

注意：
- defaultPreload: 'intent'
- scrollRestoration: true
- 环境判断用 import.meta.env.PROD，不用 process.env
```

---

## 七、组件 Prompt

```
目标：创建迁移后的基础 UI 组件。

要求：
- 全部使用 SCSS Module（`*.module.scss`），禁止内联 style
- 引用 design tokens（CSS Variables）
- 使用 clsx 包裹所有 className
- 类型定义完整，支持 forwardRef

组件清单：
- Modal
- Input（forwardRef）
- Select（forwardRef）+ options
- Textarea（forwardRef）
- Tooltip（data-position 控制方向）
- ThemePicker（主题切换 + accent 选择）
- Logo（SVG 组件）
- Form（react-hook-form + ZodSchema）
- FormField（泛型 field wrapper）

命名约定：
- 文件：kebab-case（modal.module.scss）
- selector：kebab-case
- TSX import：import classes from './modal.module.scss'
- 使用：clsx(classes.xxx)
```

---

## 八、Hooks Prompt

```
目标：迁移通用 hooks。

要求：
- 优先复用已有实现
- 保持类型安全

迁移清单（参考 [参考项目]）：
- useCopy（依赖 react-use）
- useMobile（移动端检测 + resize 监听）
- useBoolean（toggle / setTrue / setFalse）
- useDisclosure（open / close / toggle）
- useDebounce（lodash-es debounce）
```

---

## 九、Utils / Lib Prompt

```
目标：迁移通用工具函数。

要求：
- 优先复用已有实现
- 保持模块化

迁移清单：
- utils/cn.ts（clsx 导出）
- utils/date.ts（dayjs + timezone）
- utils/str.ts（truncateStr / randomStr / removeHtml）
- utils/helper.ts（openUrl / sleep）
- utils/download.ts（downloadText）
- utils/object.ts（deleteProps）
- api/request.ts（ky HTTP client）
- lib/toast.ts（showSuccess / showError / showToast）
- lib/form.ts（showFormError）
```

---

## 十、验证 Prompt

```
验证步骤（按顺序）：

1. pnpm run build（tsgo + vite build，确保无 error）
2. pnpm run check（oxlint，确保无 warning）
3. pnpm run fmt（oxfmt，格式通过）
4. pnpm run dev 启动确认无运行时错误

每完成一个步骤都汇报结果，如果有错误先修再继续。
```

---

## 常见陷阱总结

| 问题                      | 说明                           | 解决方案                                     |
| ------------------------- | ------------------------------ | -------------------------------------------- |
| `prefixUrl` → `prefix`    | ky v1 改名                     | 使用 `prefix` 而非 `prefixUrl`               |
| `process.env` 不可用      | Vite 项目                      | 改用 `import.meta.env`                       |
| SCSS 找不到 mixin         | `additionalData` 路径不对      | 确保 `@/styles/mixins` 存在且正确            |
| Zod 类型不兼容            | `zodResolver` 泛型             | 用 `as any` 绕过，类型由 schema 保证         |
| `routeTree.gen.ts` 不存在 | 自动生成文件缺失               | 手动创建或运行一次 dev 生成                  |
| 内联 style                | 不利于维护和主题切换           | 全部移到 `.module.scss` 的 class 中          |
| CSS Module 类名           | kebab-case 选择器              | 用 `camelCase` 访问（Vite localsConvention） |
| 环境变量前缀              | Vite 只暴露 `VITE_` 开头的变量 | 确保 `.env` 变量以 `VITE_` 开头              |

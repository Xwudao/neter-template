import { Link, type LinkProps } from '@tanstack/react-router'
import clsx from 'clsx'
import { type ButtonHTMLAttributes, type ReactNode } from 'react'

import classes from './button.module.scss'

type Variant = 'primary' | 'secondary' | 'ghost' | 'outline' | 'danger'
type Size = 'sm' | 'md' | 'lg'

interface BaseButtonProps {
  variant?: Variant
  size?: Size
  icon?: ReactNode
  iconRight?: ReactNode
  full?: boolean
  className?: string
  children?: ReactNode
}

type ButtonAsButton = BaseButtonProps &
  Omit<ButtonHTMLAttributes<HTMLButtonElement>, 'children'> & {
    to?: never
    href?: never
  }

type ButtonAsLink = BaseButtonProps &
  Omit<LinkProps, 'children'> & {
    to: string
    href?: never
  }

type ButtonAsAnchor = BaseButtonProps &
  Omit<React.AnchorHTMLAttributes<HTMLAnchorElement>, 'children'> & {
    to?: never
    href: string
  }

type ButtonProps = ButtonAsButton | ButtonAsLink | ButtonAsAnchor

type ButtonInnerProps = Pick<BaseButtonProps, 'icon' | 'iconRight' | 'children'>

function ButtonInner({ icon, iconRight, children }: ButtonInnerProps) {
  return (
    <>
      {icon && <span className={clsx(classes.icon, classes.iconLeft)} aria-hidden="true">{icon}</span>}
      {children && <span className={classes.label}>{children}</span>}
      {iconRight && <span className={clsx(classes.icon, classes.iconRight)} aria-hidden="true">{iconRight}</span>}
    </>
  )
}

function resolveClasses(props: BaseButtonProps): string {
  const { variant = 'secondary', size = 'md', full, className } = props
  return clsx(
    classes.button,
    classes[variant],
    classes[size],
    full && classes.full,
    className,
  )
}

/**
 * Button — renders a native `<button>`, a TanStack `<Link>`, or an `<a>` based
 * on whether `to` or `href` is provided.
 */
export function Button(props: ButtonProps) {
  const { variant = 'secondary', size = 'md', icon, iconRight, full, className, children } = props
  const innerProps = { variant, size, icon, iconRight, full, className, children }

  if ('to' in props && props.to !== undefined) {
    const { to, ...rest } = props
    return (
      <Link to={to} className={resolveClasses(innerProps)} {...(rest as Omit<LinkProps, 'to'>)}>
        <ButtonInner {...innerProps} />
      </Link>
    )
  }

  if ('href' in props && props.href !== undefined) {
    const { href, ...rest } = props
    return (
      <a href={href} className={resolveClasses(innerProps)} {...(rest as React.AnchorHTMLAttributes<HTMLAnchorElement>)}>
        <ButtonInner {...innerProps} />
      </a>
    )
  }

  const { to: _to, href: _href, variant: _variant, size: _size, icon: _icon, iconRight: _iconRight, full: _full, className: _className, children: _children, ...rest } = props as ButtonAsButton
  void _to
  void _href
  void _variant
  void _size
  void _iconRight
  void _full
  return (
    <button className={resolveClasses(innerProps)} {...rest}>
      <ButtonInner {...innerProps} />
    </button>
  )
}

export default Button

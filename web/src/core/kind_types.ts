enum KindTypes {
  BanUA = 'ban_ua',
  FriendLink = 'friend_link',
  BackupLink = 'backup_link',
}

const KindTypeLabels: Record<KindTypes, string> = {
  [KindTypes.BanUA]: '禁止UA',
  [KindTypes.FriendLink]: '友情链接',
  [KindTypes.BackupLink]: '网站备份链接',
};

interface FriendLinkValue {
  name: string;
  link: string;
  open_blank: boolean;
  enable: boolean;
}

interface BanUAValue {
  enable: boolean;
  values: string[];
}

export default KindTypes;
export { KindTypeLabels };
export type { FriendLinkValue, BanUAValue };

export type KindValueMap = {
  BanUA: BanUAValue;
  FriendLink: FriendLinkValue;
  BackupLink: FriendLinkValue; // Reusing FriendLinkValue for BackupLink
};

export type KindType = keyof typeof KindTypes;

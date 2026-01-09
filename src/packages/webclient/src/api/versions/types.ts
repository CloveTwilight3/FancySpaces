export interface SpaceVersion {
  space_id: string;
  id: string;
  name: string;
  platform: string;
  channel: string;
  published_at: Date;
  changelog: string;
  supported_platform_versions: string[];
  files: SpaceVersionFile[];
}

export interface SpaceVersionFile {
  name: string;
  url: string;
  size: number;
}

export function mapPlatformToDisplayname(name?: string): string {
  if (!name) return 'Unknown';

  switch (name.toLowerCase()) {
    case 'minecraft_plugin':
      return 'Minecraft Plugin';
    case 'minecraft_mod':
      return 'Minecraft Mod';
    case 'hytale_plugin':
      return 'Hytale Plugin';
    case 'executable':
      return 'Executable';

    default:
      return name;
  }
}

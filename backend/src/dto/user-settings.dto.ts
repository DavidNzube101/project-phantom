import { IsString, IsObject, IsOptional } from 'class-validator';

export class SaveUserSettingsDto {
  @IsString()
  walletAddress: string;

  @IsObject()
  @IsOptional()
  settings: Record<string, any>;
}

export class GetUserSettingsDto {
  @IsString()
  walletAddress: string;
}

export class UserSettingsResponseDto {
  walletAddress: string;
  settings: Record<string, any>;
  createdAt: Date;
  updatedAt: Date;
} 
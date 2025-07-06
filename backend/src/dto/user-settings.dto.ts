import { IsString, IsObject, IsOptional } from 'class-validator';
import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';

export class SaveUserSettingsDto {
  @ApiProperty({
    description: 'User wallet address (unique identifier)',
    example: '0x1234567890abcdef1234567890abcdef12345678',
    minLength: 42,
    maxLength: 42
  })
  @IsString()
  walletAddress: string;

  @ApiPropertyOptional({
    description: 'Flexible settings object that can contain any user preferences',
    example: {
      theme: 'dark',
      notifications: true,
      language: 'en',
      privacyLevel: 'high',
      autoBlockTrackers: true
    },
    additionalProperties: true
  })
  @IsObject()
  @IsOptional()
  settings: Record<string, any>;
}

export class GetUserSettingsDto {
  @ApiProperty({
    description: 'User wallet address to retrieve settings for',
    example: '0x1234567890abcdef1234567890abcdef12345678'
  })
  @IsString()
  walletAddress: string;
}

export class UserSettingsResponseDto {
  @ApiProperty({
    description: 'User wallet address',
    example: '0x1234567890abcdef1234567890abcdef12345678'
  })
  walletAddress: string;

  @ApiProperty({
    description: 'User settings object',
    example: {
      theme: 'dark',
      notifications: true,
      language: 'en',
      privacyLevel: 'high',
      autoBlockTrackers: true
    },
    additionalProperties: true
  })
  settings: Record<string, any>;

  @ApiProperty({
    description: 'Timestamp when user settings were first created',
    example: '2024-01-15T10:30:00.000Z'
  })
  createdAt: Date;

  @ApiProperty({
    description: 'Timestamp when user settings were last updated',
    example: '2024-01-15T10:30:00.000Z'
  })
  updatedAt: Date;
}

export class UpdateUserSettingsDto {
  @ApiProperty({
    description: 'Settings object to update (will be merged with existing settings)',
    example: {
      theme: 'light',
      newSetting: 'value'
    },
    additionalProperties: true
  })
  @IsObject()
  settings: Record<string, any>;
} 
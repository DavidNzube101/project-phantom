import { Injectable, NotFoundException } from '@nestjs/common';
import { SaveUserSettingsDto, GetUserSettingsDto, UserSettingsResponseDto } from '../dto/user-settings.dto';

@Injectable()
export class UserSettingsMockService {
  private readonly userSettings = new Map<string, UserSettingsResponseDto>();

  async saveUserSettings(saveUserSettingsDto: SaveUserSettingsDto): Promise<UserSettingsResponseDto> {
    const { walletAddress, settings } = saveUserSettingsDto;
    
    const now = new Date();
    const userSettingsData = {
      walletAddress,
      settings: settings || {},
      createdAt: now,
      updatedAt: now,
    };

    // Check if user already exists
    const existingUser = this.userSettings.get(walletAddress);
    
    if (existingUser) {
      // Update existing user settings
      const updatedData = {
        ...userSettingsData,
        createdAt: existingUser.createdAt, // Preserve original creation date
      };
      
      this.userSettings.set(walletAddress, updatedData);
      return updatedData;
    } else {
      // Create new user settings
      this.userSettings.set(walletAddress, userSettingsData);
      return userSettingsData;
    }
  }

  async getUserSettings(getUserSettingsDto: GetUserSettingsDto): Promise<UserSettingsResponseDto> {
    const { walletAddress } = getUserSettingsDto;
    
    const userSettings = this.userSettings.get(walletAddress);
    
    if (!userSettings) {
      throw new NotFoundException(`User settings not found for wallet address: ${walletAddress}`);
    }
    
    return userSettings;
  }

  async getUserSettingsByWalletAddress(walletAddress: string): Promise<UserSettingsResponseDto | null> {
    return this.userSettings.get(walletAddress) || null;
  }

  async updateUserSettings(walletAddress: string, settings: Record<string, any>): Promise<UserSettingsResponseDto> {
    const existingUser = this.userSettings.get(walletAddress);
    
    if (!existingUser) {
      throw new NotFoundException(`User settings not found for wallet address: ${walletAddress}`);
    }

    const updatedData = {
      ...existingUser,
      settings: { ...existingUser.settings, ...settings },
      updatedAt: new Date(),
    };

    this.userSettings.set(walletAddress, updatedData);
    return updatedData;
  }

  async deleteUserSettings(walletAddress: string): Promise<void> {
    const existingUser = this.userSettings.get(walletAddress);
    
    if (!existingUser) {
      throw new NotFoundException(`User settings not found for wallet address: ${walletAddress}`);
    }

    this.userSettings.delete(walletAddress);
  }

  // Helper method to get all settings (for debugging)
  getAllSettings(): UserSettingsResponseDto[] {
    return Array.from(this.userSettings.values());
  }

  // Helper method to clear all settings (for testing)
  clearAllSettings(): void {
    this.userSettings.clear();
  }
} 
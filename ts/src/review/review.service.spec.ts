import { Injectable } from '@nestjs/common';
import { PrismaService } from 'prisma/prisma.service';
import { Review, Prisma } from '@prisma/client';

@Injectable()
export class ReviewService {
  constructor(private prisma: PrismaService) {}

  async createReview(data: Prisma.ReviewCreateInput): Promise<Review> {
    return this.prisma.review.create({
      data,
    });
  }

  async updateReview(
    id: number,
    data: Prisma.ReviewUpdateInput,
  ): Promise<Review> {
    return this.prisma.review.update({
      where: { id },
      data,
    });
  }

  async deleteReview(id: number): Promise<Review> {
    return this.prisma.review.delete({
      where: { id },
    });
  }

  async getReviews(): Promise<Review[]> {
    return this.prisma.review.findMany();
  }
}

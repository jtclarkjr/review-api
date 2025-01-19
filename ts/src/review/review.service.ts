import { Injectable, BadRequestException } from '@nestjs/common';
import { PrismaService } from 'prisma/prisma.service';
import { Review } from '@prisma/client';
import { ReviewDto } from './review.dto';

@Injectable()
export class ReviewService {
  constructor(private prisma: PrismaService) {}

  async createReview(data: {
    review: string;
    employeeId: string;
  }): Promise<Review> {
    return this.prisma.review.create({
      data: {
        review: data.review,
        employee: {
          connect: { id: data.employeeId },
        },
      },
    });
  }

  async updateReview(
    id: string,
    data: { review?: string; employeeId?: string },
  ): Promise<Review> {
    const existingReview = await this.prisma.review.findUnique({
      where: { id },
    });

    if (!existingReview) {
      throw new BadRequestException('Review not found.');
    }

    return this.prisma.review.update({
      where: { id },
      data,
    });
  }

  async deleteReview(id: string): Promise<Review> {
    return this.prisma.review.delete({
      where: { id },
    });
  }

  async getReviews(): Promise<ReviewDto[]> {
    const reviews = await this.prisma.review.findMany({
      include: {
        reviewers: {
          include: {
            reviewer: true,
          },
        },
      },
    });

    return reviews.map((review) => ({
      id: review.id,
      employeeId: review.employeeId,
      review: review.review,
      reviewers: review.reviewers.map((assignment) => ({
        id: assignment.reviewer.id,
        name: assignment.reviewer.name,
        email: assignment.reviewer.email,
        isAdmin: assignment.reviewer.isAdmin,
      })),
    }));
  }

  async assignReviewers(
    reviewId: string,
    reviewerIds: string[],
  ): Promise<void> {
    reviewId = reviewId.toString();

    const existingAssignments = await this.prisma.reviewAssignment.findMany({
      where: {
        reviewId,
        reviewerId: { in: reviewerIds },
      },
    });

    if (existingAssignments.length > 0) {
      const existingReviewerIds = existingAssignments.map((a) => a.reviewerId);
      throw new BadRequestException(
        `Reviewer(s) with ID(s) ${existingReviewerIds.join(', ')} are already assigned to this review.`,
      );
    }

    const assignments = reviewerIds.map((reviewerId) => ({
      reviewId,
      reviewerId,
    }));

    if (assignments.length > 0) {
      await this.prisma.reviewAssignment.createMany({
        data: assignments,
      });
    }
  }
}

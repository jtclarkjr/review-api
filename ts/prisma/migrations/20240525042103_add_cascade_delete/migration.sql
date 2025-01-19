-- DropForeignKey
ALTER TABLE "Feedback" DROP CONSTRAINT "Feedback_reviewId_fkey";

-- DropForeignKey
ALTER TABLE "ReviewAssignment" DROP CONSTRAINT "ReviewAssignment_reviewId_fkey";

-- AddForeignKey
ALTER TABLE "Feedback" ADD CONSTRAINT "Feedback_reviewId_fkey" FOREIGN KEY ("reviewId") REFERENCES "Review"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "ReviewAssignment" ADD CONSTRAINT "ReviewAssignment_reviewId_fkey" FOREIGN KEY ("reviewId") REFERENCES "Review"("id") ON DELETE CASCADE ON UPDATE CASCADE;

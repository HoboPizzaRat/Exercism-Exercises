using System.Linq.Expressions;

public static class Darts
{
    public static int Score(double x, double y)
    {
       double distanceFromCenter = Math.Sqrt(Math.Pow(x, 2) + Math.Pow(y, 2));

       if(distanceFromCenter <= 1)
        {
            return 10;
        }
        else if(distanceFromCenter <= 5)
        {
            return 5;
        }
        else if(distanceFromCenter <= 10)
        {
            return 1;
        }
        return 0;
    }
}
